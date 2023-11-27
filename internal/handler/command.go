package handler

import (
	"fmt"
	"tg_bot/config"
	"tg_bot/internal/common"
	"tg_bot/internal/logger"
	"tg_bot/internal/models"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleCommand(update tgbotapi.Update, bot models.BotAPI, userData *models.UserState) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	command := update.Message.Command()

	logger.Debug(fmt.Sprintf("[%s:%s] command: %s", userData.UserName, userData.Action, command))

	switch command {
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
	userData.Action = command
}
