package app

import (
	"os"
	"sync"
	"tg_bot/internal/common"
	"tg_bot/internal/database"
	"tg_bot/internal/handler"
	"tg_bot/internal/logger"
	"tg_bot/internal/otrs"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	userStates sync.Map
	err        error
)

func setupBotAPI() *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(common.GetBotToken())
	if err != nil {
		// logger.Error(err.Error())
	}
	bot.Debug = false
	return bot
}

func run(bot *tgbotapi.BotAPI, logger *logger.Logger) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	// logger.Info("Bot started")
	for update := range updates {
		handler.Run(update, bot, &userStates, logger)
	}
}

func Run() {
	database.ConnectDB()
	otrs.SetupOTRSConnector()
	bot := setupBotAPI()
	logger := logger.New(os.Getenv("DEBUG") == "1")

	run(bot, logger)
}
