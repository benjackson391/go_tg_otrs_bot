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
var GetTicketArticles = GetTicketArticlesImpl

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

type Article struct {
	SenderTypeName string `gorm:"column:sender_type_name" json:"sender_type_name"`
	ChannelName    string `gorm:"column:channel_name" json:"channel_name"`
	From           string `gorm:"column:a_from" json:"a_from"`
	Subject        string `gorm:"column:subject" json:"subject"`
	Body           string `gorm:"column:body" json:"body"`
	CreateTime     string `gorm:"column:create_time" json:"create_time"`
}

func GetTicketArticlesImpl(TicketID string) []Article {
	var articles []Article

	_ = db.Raw(`
		SELECT
			sender_type.name AS sender_type_name,
			channel.name AS channel_name,
			plain.a_from AS a_from,
			plain.a_subject as subject,
			plain.a_body AS body,
			a.create_time as create_time
		FROM article AS a
		JOIN article_sender_type AS sender_type ON a.article_sender_type_id = sender_type.id
		JOIN communication_channel AS channel ON a.communication_channel_id = channel.id
		JOIN article_data_mime AS plain ON a.id = plain.article_id
		WHERE a.is_visible_for_customer = 1 AND
			a.ticket_id = ?
	`, TicketID).Scan(&articles).Error

	return articles
}
