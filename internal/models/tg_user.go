package models

type TgUser struct {
	// gorm.Model
	TelegramLogin     string `gorm:"column:TelegramLogin"`
	CustomerUserLogin string `gorm:"column:CustomerUserLogin"`
}
