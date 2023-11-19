package handler

import (
	"tg_bot/internal/models"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/newrelic/go-agent/v3/newrelic"

	"github.com/stretchr/testify/mock"
)

var (
	userData *models.UserState
	txn      *newrelic.Transaction
)

type BotAPI interface {
	Send(c tgbotapi.Chattable) (tgbotapi.Message, error)
	GetFile(c tgbotapi.FileConfig) (tgbotapi.File, error)
	GetFileDirectURL(fileID string) (string, error)
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

func (m *MockBotAPI) GetFileDirectURL(fileID string) (string, error) {
	args := m.Called(fileID)
	return "https://api.telegram.org/file/bot/", args.Error(1)
}
