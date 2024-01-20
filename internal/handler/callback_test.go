package handler

import (
	"fmt"
	"testing"
	"tg_bot/config"
	"tg_bot/internal/logger"
	"tg_bot/internal/models"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandleCallbackQuery(t *testing.T) {
	bot := new(MockBotAPI)

	logger := logger.New(true)

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
		HandleCallbackQuery(update, bot, test_case.UserData, logger)

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
