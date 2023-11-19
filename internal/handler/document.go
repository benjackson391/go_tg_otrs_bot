package handler

import (
	"tg_bot/config"
	"tg_bot/internal/common"
	"tg_bot/internal/models"
	"tg_bot/internal/otrs"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func HandleDocument(update tgbotapi.Update, bot models.BotAPI, userData *models.UserState) {
	doc := update.Message.Document

	if doc.FileSize > config.FileSizeLimit {
		bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, config.BigFileMessage))
	} else {
		documentBytes, err := common.GetFileContent(doc, bot)
		if err != nil {
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, config.ErrorMsq1))
		} else {
			_, text := otrs.CreateOrUpdateTicket(bot, userData, doc, &documentBytes)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
			buttons := []models.Button{
				{Title: "Назад", Callback: "start"},
			}
			msg.ReplyMarkup = common.GetInlineKeyboard(buttons)
			bot.Send(msg)
			common.CleanUpUserData(userData)
		}
	}

}
