package models

type UserState struct {
	UserName          string
	Topic             string
	Description       string
	Email             string
	Code              int
	Action            string
	TicketID          string
	CustomerUserLogin string
	Vote              *string
}
