package models

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type BotAPI interface {
	Send(c tgbotapi.Chattable) (tgbotapi.Message, error)
	GetFile(c tgbotapi.FileConfig) (tgbotapi.File, error)
	GetFileDirectURL(fileID string) (string, error)
}
