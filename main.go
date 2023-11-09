package main

// 1.auth
// 1.1 get_email
// 1.2 confirm_code
// 2.main_screen
// 2.1 new_ticket
// 2.1.1 get_title
// 2.1.2 get_description
// 2.1.3 get_attachment
// 2.1.4 create_ticket
// 2.2		"Проверить статус"		checkStatus()
// 2.2.1 chow state_type_count
// 2.3 update_request

import (
	"context"
	"os"
	"sync"

	"tg_bot/app"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/withmandala/go-log"
)

var (
	botToken       string
	err            error
	newrelic_token string
	Logger         *log.Logger
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		Logger.Fatal("Error loading .env file")
	}
	botToken = os.Getenv("BOT_TOKEN")
	newrelic_token = os.Getenv("NEWRELIC")
}

func setupLogger() {
	Logger = log.New(os.Stderr).WithColor().WithDebug()
	app.InitLogger(Logger)
}

func setupBotAPI() *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		Logger.Fatal(err)
	}
	bot.Debug = false
	return bot
}

func run(bot *tgbotapi.BotAPI) {
	appNR, err := newrelic.NewApplication(
		newrelic.ConfigAppName("tg_bot"),
		newrelic.ConfigLicense("eu01xx77231d3545dab7a73708119005FFFFNRAL"),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)
	if err != nil {
		Logger.Fatal(err)
	}

	ctx := context.Background()
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, _ := bot.GetUpdatesChan(u)

	Logger.Info("Bot started")

	var wg sync.WaitGroup

	for {
		select {
		case update := <-updates:
			txn := appNR.StartTransaction("run")
			defer txn.End()

			wg.Add(1)
			go app.Handle(&wg, update, bot, txn)
		case <-ctx.Done():
			Logger.Info("Bot stopped due to context cancellation")
			return
		}
	}
}

func main() {
	loadEnv()
	setupLogger()
	app.ConnectDB()
	app.SetupOTRSConnector()
	bot := setupBotAPI()
	run(bot)
}
