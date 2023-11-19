package app

import (
	"os"
	"runtime/pprof"
	"sync"

	"tg_bot/internal/database"
	"tg_bot/internal/handler"
	"tg_bot/internal/logger"
	"tg_bot/internal/otrs"

	"golang.org/x/net/context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/newrelic/go-agent/v3/newrelic"
)

var err error

func setupBotAPI() *tgbotapi.BotAPI {
	botToken := os.Getenv("BOT_TOKEN")
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		logger.Error(err.Error())
	}
	bot.Debug = false
	return bot
}

func run(bot *tgbotapi.BotAPI) {
	cpuFile, err := os.Create("cpu.prof")
	if err != nil {
		panic(err)
	}
	defer cpuFile.Close()

	// Запуск профилирования CPU
	if err := pprof.StartCPUProfile(cpuFile); err != nil {
		panic(err)
	}
	defer pprof.StopCPUProfile()

	memFile, err := os.Create("mem.prof")
	if err != nil {
		panic(err)
	}
	defer memFile.Close()
	if err := pprof.WriteHeapProfile(memFile); err != nil {
		panic(err)
	}

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

	updates, _ := bot.GetUpdatesChan(u)

	logger.Info("Bot started")

	var wg sync.WaitGroup

	for {
		select {
		case update := <-updates:
			txn := appNR.StartTransaction("run")
			defer txn.End()

			wg.Add(1)
			go handler.Run(&wg, update, bot)
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
