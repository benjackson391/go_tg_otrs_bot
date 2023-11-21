package handler

import (
	"tg_bot/config"
	"tg_bot/internal/common"
	"tg_bot/internal/logger"
	"tg_bot/internal/models"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func HandleCommand(update tgbotapi.Update, bot models.BotAPI, userData *models.UserState) {
	logger.Debug("command.HandleCommand")
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	switch update.Message.Command() {
	case "start":
		msg.Text = config.WelcomeMessage + "\n\n" + config.WelcomeDescription
		msg.ReplyMarkup = common.GetInitialKeyboard()
		bot.Send(msg)
	case "stop":
		msg.Text = config.LeaveMessage
		bot.Send(msg)
		common.CleanUpUserData(userData)
	default:
		msg.Text = config.UnknownCommand
		bot.Send(msg)
	}
}
