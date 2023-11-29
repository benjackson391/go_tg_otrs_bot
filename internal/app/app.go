package app

import (
	"os"
	"sync"

	"tg_bot/internal/database"
	"tg_bot/internal/handler"
	"tg_bot/internal/logger"
	"tg_bot/internal/otrs"

	"golang.org/x/net/context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/newrelic/go-agent/v3/newrelic"
)

var err error

func setupBotAPI() *tgbotapi.BotAPI {
	botToken := os.Getenv("PROD_BOT_TOKEN")
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		logger.Error(err.Error())
	}
	bot.Debug = false
	return bot
}

func run(bot *tgbotapi.BotAPI) {
	appNR, err := newrelic.NewApplication(
		newrelic.ConfigAppName("tg_bot"),
		newrelic.ConfigLicense(os.Getenv("NEWRELIC")),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)
	if err != nil {
		logger.Error(err.Error())
	}

	ctx := context.Background()
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	logger.Info("Bot started")

	var wg sync.WaitGroup
	var userStates sync.Map

	for {
		select {
		case update := <-updates:
			txn := appNR.StartTransaction("run")
			defer txn.End()

			wg.Add(1)
			go handler.Run(&wg, update, bot, &userStates)
		case <-ctx.Done():
			logger.Info("Bot stopped due to context cancellation")
			return
		}
	}
}

func Run() {
	if os.Getenv("DEBUG") != "" {
		logger.EnableDebug()
	}
	database.ConnectDB()
	otrs.SetupOTRSConnector()
	bot := setupBotAPI()
	run(bot)
}
