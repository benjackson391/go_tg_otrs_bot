package handler

import (
	"fmt"
	"regexp"
	"strings"
	"sync"
	"tg_bot/internal/common"
	"tg_bot/internal/database"
	"tg_bot/internal/logger"
	"tg_bot/internal/models"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	ticketRegex = regexp.MustCompile(`^\d+$`)
	voteRegex   = regexp.MustCompile(`^vote_?(\d)?$`)
)

func HandleUpdate(update tgbotapi.Update, bot models.BotAPI, userData *models.UserState, logger *logger.Logger) {
	userData.Trace = append(userData.Trace, "HandleUpdate")

	if update.Message != nil {
		if update.Message.IsCommand() {
			HandleCommand(update, bot, userData, logger)
		} else {
			HandleMessage(update, bot, userData, logger)
		}
	} else if update.CallbackQuery != nil {
		HandleCallbackQuery(update, bot, userData, logger)
	}
}

func Run(update tgbotapi.Update, bot models.BotAPI, userStates *sync.Map, logger *logger.Logger) {
	userID := common.GetUserID(update)
	userName := common.GetUserName(update)
	userData := common.GetUserData(userID, userName, userStates)

	isAuthorized, CustomerUserLogin := database.IsAuthorized(userName)
	userData.CustomerUserLogin = CustomerUserLogin

	stateBefore := userData.State
	userData.Trace = []string{"Run"}

	if isAuthorized {
		HandleUpdate(update, bot, &userData, logger)
	} else {
		handleAuthoriseMessage(update, bot, &userData, logger)
	}
	userStates.Store(userID, userData)

	stateAfter := userData.State
	logger.Info("update", "UserName", userData.UserName, "state", fmt.Sprintf("%d -> %d", stateBefore, stateAfter), "trace", strings.Join(userData.Trace, "::"))
}
