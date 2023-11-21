package handler

import (
	"testing"
	"tg_bot/config"
	"tg_bot/internal/models"

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
