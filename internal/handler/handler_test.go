package handler

import (
	"fmt"
	"testing"
	"tg_bot/config"
	"tg_bot/internal/models"

	"github.com/jarcoal/httpmock"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

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
	bot := new(MockBotAPI)

	update := getUpdateObj("/start")

	bot.On("Send", mock.Anything).Return(tgbotapi.Message{}, nil)

	HandleCommand(update, bot, userData)
	if len(bot.SentMessages) != 1 {
		t.Errorf("Expected one message to be sent for start command, got: %v", len(bot.SentMessages))
	} else {
		answer, ok := bot.SentMessages[0].(tgbotapi.MessageConfig)
		if !ok {
			t.Errorf("Expected message to be of type tgbotapi.MessageConfig")
		} else {
			assert.Equal(t, config.WelcomeMessage+"\n\n"+config.WelcomeDescription, answer.Text, "Text should be equal")

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
	bot := new(MockBotAPI)

	update := getUpdateObj("/stop")

	bot.On("Send", mock.Anything).Return(tgbotapi.Message{}, nil)

	userData = &models.UserState{
		CurrentState: "test",
	}

	HandleCommand(update, bot, userData)
	if len(bot.SentMessages) != 1 {
		t.Errorf("Expected one message to be sent for start command, got: %v", len(bot.SentMessages))
	} else {
		answer, ok := bot.SentMessages[0].(tgbotapi.MessageConfig)
		if !ok {
			t.Errorf("Expected message to be of type tgbotapi.MessageConfig")
		} else {
			assert.Equal(t, config.LeaveMessage, answer.Text, "Text should be equal")
		}
	}
}

func TestHandleCommandDefault(t *testing.T) {
	bot := new(MockBotAPI)

	update := getUpdateObj("/unknown")

	bot.On("Send", mock.Anything).Return(tgbotapi.Message{}, nil)

	HandleCommand(update, bot, userData)
	if len(bot.SentMessages) != 1 {
		t.Errorf("Expected one message to be sent for start command, got: %v", len(bot.SentMessages))
	} else {
		answer, ok := bot.SentMessages[0].(tgbotapi.MessageConfig)
		if !ok {
			t.Errorf("Expected message to be of type tgbotapi.MessageConfig")
		} else {
			assert.Equal(t, config.UnknownCommand, answer.Text, "Text should be equal")
		}
	}
}

func TestHandleMessage(t *testing.T) {

}

func TestHandleCallbackQuery(t *testing.T) {
	bot := new(MockBotAPI)

	csenario := []struct {
		UserData     *models.UserState
		CallbackData string
		AnswerText   string
	}{
		{
			UserData:     &models.UserState{},
			CallbackData: "start",
			AnswerText:   config.WelcomeDescription,
		},
		{
			UserData:     &models.UserState{},
			CallbackData: "vote_5",
			AnswerText:   "Ошибка: недостаточно данных, чтобы создать/обновить заявку!",
		},
		{
			UserData: &models.UserState{
				TicketID: "1",
			},
			CallbackData: "vote_5",
			AnswerText:   "Обращение 1 оценено!",
		},
		{
			UserData: &models.UserState{
				TicketID: "1",
			},
			CallbackData: "vote_4",
			AnswerText:   "Обращение 1 оценено!",
		},
		{
			UserData: &models.UserState{
				TicketID: "1",
			},
			CallbackData: "vote_3",
			AnswerText:   "Обращение 1 оценено!",
		},
		{
			UserData: &models.UserState{
				TicketID: "1",
			},
			CallbackData: "vote_2",
			AnswerText:   "Обращение 1 оценено!",
		},
		{
			UserData: &models.UserState{
				TicketID: "1",
			},
			CallbackData: "vote_",
			AnswerText:   "Обращение 1 оценено!",
		},
	}

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	for test_number, test_case := range csenario {
		update := tgbotapi.Update{
			CallbackQuery: &tgbotapi.CallbackQuery{
				Data: test_case.CallbackData,
				Message: &tgbotapi.Message{
					Chat: &tgbotapi.Chat{
						ID: 12345,
					},
				},
			},
		}

		httpmock.RegisterResponder("POST", "/create",
			httpmock.NewStringResponder(200, `{"ArticleID":"1","TicketID":"1","TicketNumber":"1"}`))
		httpmock.RegisterResponder("POST", "/update",
			httpmock.NewStringResponder(200, `{"ArticleID":"1","TicketID":"1","TicketNumber":"1"}`))

		bot.On("Send", mock.Anything).Return(tgbotapi.Message{}, nil)
		HandleCallbackQuery(update, bot, test_case.UserData)

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
