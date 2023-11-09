package models

type TgTicket struct {
	// gorm.Model
	TicketID                    int64   `gorm:"column:TicketID"`
	TicketNumber                string  `gorm:"column:TicketNumber"` // если вы используете ORM, то db теги указываются в соответствии с названиями полей в таблице
	Title                       *string `gorm:"column:Title"`        // указываем как указатель, потому что может быть NULL
	State                       string  `gorm:"column:State"`
	TelegramLogin               *string `gorm:"column:TelegramLogin"` // может быть NULL, поэтому используем указатель
	Type                        string  `gorm:"column:Type"`
	StateType                   string  `gorm:"column:StateType"`
	UserFirstname               string  `gorm:"column:UserFirstname"`
	UserLastname                string  `gorm:"column:UserLastname"`
	CustomerUserLogin           string  `gorm:"column:CustomerUserLogin"`
	Body                        *string `gorm:"column:Body"` // может быть NULL
	SolutionTimeDestinationDate int     `gorm:"column:SolutionTimeDestinationDate"`
	Timestamp                   int64   `gorm:"column:timestamp"`
	TimestampNow                int64   `gorm:"column:timestamp_now"`
}
