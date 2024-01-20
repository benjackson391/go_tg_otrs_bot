package handler

import (
	"tg_bot/config"
	"tg_bot/internal/common"
	"tg_bot/internal/logger"
	"tg_bot/internal/models"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleCommand(update tgbotapi.Update, bot models.BotAPI, userData *models.UserState, logger *logger.Logger) {
	userData.Trace = append(userData.Trace, "HandleCommand")
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	command := update.Message.Command()

	switch command {
	case "start":
		userData.State = 1

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
	userData.Action = command
}
