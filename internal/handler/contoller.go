package handler

import (
	"regexp"
	"sync"
	"tg_bot/internal/common"
	"tg_bot/internal/database"
	"tg_bot/internal/logger"
	"tg_bot/internal/models"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	ticketRegex = regexp.MustCompile(`^\d+$`)
	voteRegex   = regexp.MustCompile(`^vote_?(\d)?$`)
)

func HandleUpdate(update tgbotapi.Update, bot models.BotAPI, userData *models.UserState) {
	logger.Debug("controller.HandleUpdate")
	if update.Message != nil {
		if update.Message.IsCommand() {
			HandleCommand(update, bot, userData)
		} else {
			HandleMessage(update, bot, userData)
		}
	} else if update.CallbackQuery != nil {
		HandleCallbackQuery(update, bot, userData)
	}
}

func Run(wg *sync.WaitGroup, update tgbotapi.Update, bot models.BotAPI, userStates *sync.Map) {
	logger.Debug("controller.Run")
	defer wg.Done()

	userID := common.GetUserID(update)
	userName := common.GetUserName(update)
	userData := common.GetUserData(userID, userName, userStates)

	isAuthorized, CustomerUserLogin := database.IsAuthorized(userName)
	userData.CustomerUserLogin = CustomerUserLogin

	if isAuthorized {
		HandleUpdate(update, bot, &userData)
	} else {
		handleAuthoriseMessage(update, bot, &userData)
	}
	userStates.Store(userID, userData)
	logger.Debug("userData after:")
}
