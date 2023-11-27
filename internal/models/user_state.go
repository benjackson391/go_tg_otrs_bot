package models

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type UserState struct {
	UserName          string
	Topic             string
	Description       string
	Email             string
	Code              int
	Action            string
	TicketID          string
	CustomerUserLogin string
	Vote              *string
	Chan              chan *tgbotapi.Document
	MediaGroupID      string
}
