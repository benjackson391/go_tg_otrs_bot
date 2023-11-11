package app

import (
	"testing"
	"tg_bot/models"

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
