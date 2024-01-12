package database

import (
	"log"
	"os"
	"tg_bot/internal/database/dal"
	"tg_bot/internal/models"
	"time"

	"gorm.io/gorm/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	dsn       string
	db        *gorm.DB
	Q         *dal.Query
	newLogger logger.Interface
)

func ConnectDB() {
	dsn = os.Getenv("DB_DSN")

	newLogger = logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Пороговое время для медленных запросов
			LogLevel:      logger.Warn, // Уровень логирования
			Colorful:      true,        // Включить цвета
		},
	)

	db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
		Logger:                 newLogger,
	})

	dal.SetDefault(db)
	Q = dal.Use(db)
}

var IsAuthorized = isAuthorizedImpl
var GetTickets = GetTicketsImpl
var GetTicket = GetTicketImpl

func isAuthorizedImpl(userName string) (bool, string) {
	var user models.TgUser
	db.First(&user, "TelegramLogin = ?", userName)

	IsAuthorized := user != models.TgUser{}

	return IsAuthorized, user.CustomerUserLogin
}

func GetStateTypeCount(CustomerUserLogin string) models.StateTypeCount {
	var count models.StateTypeCount

	_ = db.Raw(`
	SELECT
	  SUM(CASE WHEN StateType = 'pending auto' THEN 1 ELSE 0 END) AS count_pend_auto,
	  SUM(CASE WHEN StateType != 'pending auto' THEN 1 ELSE 0 END) AS count_open
	FROM
	  tg_tickets
	WHERE
	  CustomerUserLogin LIKE ?
	`, CustomerUserLogin).Scan(&count).Error

	return count
}

func GetTicketsImpl(userName string) []models.TgTicket {
	var tickets []models.TgTicket

	db.Find(&tickets, "TelegramLogin = ?", userName)

	return tickets
}

func GetTicketImpl(TicketID string) models.TgTicket {
	var ticket models.TgTicket

	db.First(&ticket, "TicketID = ?", TicketID)

	return ticket
}
