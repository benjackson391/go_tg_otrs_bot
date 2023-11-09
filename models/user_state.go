package models

type UserState struct {
	Stage             string
	IsAuthorized      bool
	UserName          string
	CurrentState      string
	Topic             string
	Description       string
	Email             string
	Code              int
	Action            string
	TicketID          string
	CustomerUserLogin string
}
