package models

type UserState struct {
	UserName          string
	CurrentState      string
	Topic             string
	Description       string
	Email             string
	Code              int
	Action            string
	TicketID          string
	CustomerUserLogin string
	Vote              *string
}
