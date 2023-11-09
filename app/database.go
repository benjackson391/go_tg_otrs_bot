package app

import (
	"fmt"
	"os"
	"tg_bot/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	err     error
	dsn     string
	db      *gorm.DB
	myCache = NewCacheMap()
)

func ConnectDB() {
	dsn = os.Getenv("DB_DSN")
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		Logger.Fatal(err)
	}
}

func IsAuthorized(userName string) (bool, string) {
	var IsAuthorized bool
	// if value, found := myCache.Get(userName); found {
	// 	boolVal, err := strconv.ParseBool(value)
	// 	if err != nil {
	// 		Logger.Fatal(err) // Обработка ошибки, если входная строка не является допустимым булевым значением
	// 	}
	// 	return boolVal,
	// }

	var user models.TgUser
	db.First(&user, "TelegramLogin = ?", userName)

	IsAuthorized = user != models.TgUser{}
	// if IsAuthorized == true {
	// 	myCache.Set(userName, strconv.FormatBool(IsAuthorized), 30*time.Second)
	// }

	Logger.Debug(fmt.Sprintf("[util:IsAuthorized] user \"%s\" auth: %t", userName, IsAuthorized))
	return IsAuthorized, user.CustomerUserLogin
}

func getStateTypeCount(CustomerUserLogin string) models.StateTypeCount {
	var count models.StateTypeCount

	err = db.Raw(`
	SELECT
	  SUM(CASE WHEN StateType = 'pending auto' THEN 1 ELSE 0 END) AS count_pend_auto,
	  SUM(CASE WHEN StateType != 'pending auto' THEN 1 ELSE 0 END) AS count_open
	FROM
	  tg_tickets
	WHERE
	  CustomerUserLogin LIKE ?
	`, CustomerUserLogin).Scan(&count).Error

	if err != nil {
		Logger.Warn("Query failed", err)
	}

	return count
}

func getTickets(userName string) []models.TgTicket {
	var tickets []models.TgTicket

	db.Find(&tickets, "TelegramLogin = ?", userName)

	return tickets
}
