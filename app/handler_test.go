package app

import (
	"fmt"
	"os"
	"testing"
	"tg_bot/models"

	"github.com/withmandala/go-log"

	"github.com/jarcoal/httpmock"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	userData *models.UserState
	txn      *newrelic.Transaction
)

type BotAPI interface {
	Send(c tgbotapi.Chattable) (tgbotapi.Message, error)
	GetFile(c tgbotapi.FileConfig) (tgbotapi.File, error)
}

type MockBotAPI struct {
	mock.Mock
	SentMessages []tgbotapi.Chattable
}

func (m *MockBotAPI) Send(c tgbotapi.Chattable) (tgbotapi.Message, error) {
	args := m.Called(c)
	m.SentMessages = append(m.SentMessages, c)
	return args.Get(0).(tgbotapi.Message), args.Error(1)
}

func (m *MockBotAPI) GetFile(c tgbotapi.FileConfig) (tgbotapi.File, error) {
	args := m.Called(c)
	return args.Get(0).(tgbotapi.File), args.Error(1)
}

func getUpdateObj(command string) tgbotapi.Update {
	return tgbotapi.Update{
		Message: &tgbotapi.Message{
			Text: command,
			Chat: &tgbotapi.Chat{
				ID: 12345,
			},
			Entities: &[]tgbotapi.MessageEntity{
				{
					Type:   "bot_command",
					Offset: 0,
					Length: len(command),
				},
			},
		},
	}
}

func TestHandleCommandStart(t *testing.T) {
	InitLogger(log.New(os.Stderr).WithColor().WithDebug())
	bot := new(MockBotAPI)

	update := getUpdateObj("/start")

	bot.On("Send", mock.Anything).Return(tgbotapi.Message{}, nil)

	HandleCommand(update, bot, userData, txn)
	if len(bot.SentMessages) != 1 {
		t.Errorf("Expected one message to be sent for start command, got: %v", len(bot.SentMessages))
	} else {
		answer, ok := bot.SentMessages[0].(tgbotapi.MessageConfig)
		if !ok {
			t.Errorf("Expected message to be of type tgbotapi.MessageConfig")
		} else {
			assert.Equal(t, WelcomeMessage+"\n\n"+WelcomeDescription, answer.Text, "Text should be equal")

			expectedKeyboard := tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("Новая заявка", "new_request"),
				),
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("Проверить статус", "check_status"),
				),
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("Обновить заявку", "update_request"),
				),
			)
			assert.Equal(t, &expectedKeyboard, answer.ReplyMarkup, "NewInlineKeyboardMarkup should be equal")
		}
	}
}

func TestHandleCommandStop(t *testing.T) {
	InitLogger(log.New(os.Stderr).WithColor().WithDebug())
	bot := new(MockBotAPI)

	update := getUpdateObj("/stop")

	bot.On("Send", mock.Anything).Return(tgbotapi.Message{}, nil)

	HandleCommand(update, bot, userData, txn)
	if len(bot.SentMessages) != 1 {
		t.Errorf("Expected one message to be sent for start command, got: %v", len(bot.SentMessages))
	} else {
		answer, ok := bot.SentMessages[0].(tgbotapi.MessageConfig)
		if !ok {
			t.Errorf("Expected message to be of type tgbotapi.MessageConfig")
		} else {
			assert.Equal(t, LeaveMessage, answer.Text, "Text should be equal")
		}
	}
}

func TestHandleCommandDefault(t *testing.T) {
	InitLogger(log.New(os.Stderr).WithColor().WithDebug())
	bot := new(MockBotAPI)

	update := getUpdateObj("/unknown")

	bot.On("Send", mock.Anything).Return(tgbotapi.Message{}, nil)

	HandleCommand(update, bot, userData, txn)
	if len(bot.SentMessages) != 1 {
		t.Errorf("Expected one message to be sent for start command, got: %v", len(bot.SentMessages))
	} else {
		answer, ok := bot.SentMessages[0].(tgbotapi.MessageConfig)
		if !ok {
			t.Errorf("Expected message to be of type tgbotapi.MessageConfig")
		} else {
			assert.Equal(t, UnknownCommand, answer.Text, "Text should be equal")
		}
	}
}

func TestHandleDocument(t *testing.T) {
	InitLogger(log.New(os.Stderr).WithColor().WithDebug())
	bot := new(MockBotAPI)

	csenario := []struct {
		UserData     *models.UserState
		FileSize     int
		OtrsPath     string
		OtrsResponse string
		AnswerText   string
		Code         int
	}{
		{
			UserData:     &models.UserState{TicketID: ""},
			FileSize:     FileSizeLimit + FileSizeLimit*0.5,
			OtrsPath:     "/create",
			OtrsResponse: "",
			Code:         200,
			AnswerText:   "Файл слишком большой. Приложите другой файл",
		},
		{
			UserData:     &models.UserState{TicketID: ""},
			FileSize:     FileSizeLimit,
			OtrsPath:     "/create",
			OtrsResponse: `{"ArticleID":"1","TicketID":"1","TicketNumber":"1"}`,
			Code:         200,
			AnswerText:   fmt.Sprintf(TicketCreatedTemplate, "1"),
		},
		{
			UserData:     &models.UserState{TicketID: ""},
			FileSize:     FileSizeLimit * 0.5,
			OtrsPath:     "/create",
			OtrsResponse: `{"ArticleID":"1","TicketID":"2","TicketNumber":"2"}`,
			AnswerText:   fmt.Sprintf(TicketCreatedTemplate, "2"),
		},
		{
			UserData:     &models.UserState{TicketID: "1"},
			FileSize:     FileSizeLimit * 0.5,
			OtrsPath:     "/update",
			OtrsResponse: `{"ArticleID":"1","TicketID":"2","TicketNumber":"3"}`,
			AnswerText:   fmt.Sprintf(TicketUpdatedMessage, "3"),
		},
	}

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.telegram.org/file/bot/",
		httpmock.NewStringResponder(200, "<content>"))

	for test_number, test_case := range csenario {
		httpmock.RegisterResponder("POST", test_case.OtrsPath,
			httpmock.NewStringResponder(test_case.Code, test_case.OtrsResponse))

		document := tgbotapi.Document{
			FileID:   "file",
			FileName: "example.pdf",
			MimeType: "application/pdf",
			FileSize: test_case.FileSize,
		}

		message := &tgbotapi.Message{
			Document: &document,
			Chat:     &tgbotapi.Chat{ID: 12345},
		}
		update := tgbotapi.Update{
			Message: message,
		}

		bot.On("GetFile", mock.Anything).Return(tgbotapi.File{}, nil)
		bot.On("Send", mock.Anything).Return(tgbotapi.Message{}, nil)

		HandleDocument(update, bot, test_case.UserData, txn)

		if len(bot.SentMessages) != 1 {
			t.Errorf("[%d] Expected one message to be sent for start command, got: %v", test_number, len(bot.SentMessages))
		} else {
			answer, ok := bot.SentMessages[0].(tgbotapi.MessageConfig)
			if !ok {
				t.Errorf("[%d] Expected message to be of type tgbotapi.MessageConfig", test_number)
			} else {
				assert.Equal(t, test_case.AnswerText, answer.Text, fmt.Sprintf("[%d] Text should be equal", test_number))
			}
		}
		bot.SentMessages = []tgbotapi.Chattable{}
	}
}
