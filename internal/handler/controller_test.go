package handler

import (
	"fmt"
	"sync"
	"testing"
	"tg_bot/config"
	"tg_bot/internal/database"
	"tg_bot/internal/logger"
	"tg_bot/internal/models"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockDBService struct {
	mock.Mock
}

func (m *MockDBService) Find(dest interface{}, conds ...interface{}) *gorm.DB {
	args := m.Called(dest, conds)
	return args.Get(0).(*gorm.DB)
}

var wg sync.WaitGroup

func TestUpdateTicket(t *testing.T) {
	bot := new(MockBotAPI)

	logger.EnableDebug()

	database.IsAuthorized = func(userName string) (bool, string) {
		return true, "CustomerUserLogin"
	}

	mock1 := "mock"

	database.GetTickets = func(userName string) []models.TgTicket {
		return []models.TgTicket{
			{
				Title:        &mock1,
				TicketNumber: mock1,
			},
		}
	}

	database.GetTicket = func(TicketID string) models.TgTicket {
		return models.TgTicket{
			TicketNumber:                mock1,
			Type:                        mock1,
			UserFirstname:               "-",
			UserLastname:                "-",
			State:                       "new",
			SolutionTimeDestinationDate: 12123,
			Body:                        &mock1,
		}
	}

	update := tgbotapi.Update{
		CallbackQuery: &tgbotapi.CallbackQuery{
			Data: "start",
			Message: &tgbotapi.Message{
				Chat: &tgbotapi.Chat{
					ID: 12345,
				},
			},
			From: &tgbotapi.User{
				ID:       12345,
				UserName: "UserName",
			},
		},
	}

	var userStates sync.Map

	bot.On("Send", mock.Anything).Return(tgbotapi.Message{}, nil)
	wg.Add(1)
	Run(&wg, update, bot, &userStates)

	answer, ok := bot.SentMessages[0].(tgbotapi.MessageConfig)
	if !ok {
		t.Errorf("[%d] Expected message to be of type tgbotapi.MessageConfig", 1)
	} else {
		assert.Equal(t, config.WelcomeDescription, answer.Text, fmt.Sprintf("[%d] Text should be equal", 1))
	}

	update.CallbackQuery.Data = "update_request"
	wg.Add(1)
	Run(&wg, update, bot, &userStates)

	answer, ok = bot.SentMessages[1].(tgbotapi.MessageConfig)
	if !ok {
		t.Errorf("[%d] Expected message to be of type tgbotapi.MessageConfig", 2)
	} else {
		assert.Equal(t, config.UpdateTitle, answer.Text, fmt.Sprintf("[%d] Text should be equal", 2))
	}

	// update.CallbackQuery.Data = "1"
	// wg.Add(1)
	// Run(&wg, update, bot, &userStates)

	// answer, ok = bot.SentMessages[2].(tgbotapi.MessageConfig)
	// if !ok {
	// 	t.Errorf("[%d] Expected message to be of type tgbotapi.MessageConfig", 3)
	// } else {
	// 	assert.Equal(t, "Номер заявки:\t#mock\nTип:\tmock\nИсполнитель:\t- -\nСтатус:\tновая\nПредельное время решения:\t1970-01-01 06:22:03\n\nОписание:\tmock", answer.Text, fmt.Sprintf("[%d] Text should be equal", 3))

	// 	buttons := []models.Button{
	// 		{Title: "Добавить комментарий", Callback: "add_comment"},
	// 		{
	// 			Title:    "Назад",
	// 			Callback: "start",
	// 		},
	// 	}

	// 	assert.Equal(t, common.GetInlineKeyboard(buttons), answer.ReplyMarkup, fmt.Sprintf("[%d] ReplyMarkup should be equal", 3))
	// }

}
