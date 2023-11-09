package models

type StateTypeCount struct {
	CountPendAuto int
	CountOpen     int
}

type Article struct {
	CommunicationChannel string `json:"CommunicationChannel"`
	SenderType           string `json:"SenderType"`
	Charset              string `json:"Charset"`
	MimeType             string `json:"MimeType"`
	From                 string `json:"From"`
	Subject              string `json:"Subject"`
	Body                 string `json:"Body"`
}

type Ticket struct {
	Title        *string `json:"Title"`
	QueueID      *int    `json:"QueueID"`
	TypeID       *int    `json:"TypeID"`
	CustomerUser *string `json:"CustomerUser"`
	StateID      *int    `json:"StateID"`
	PriorityID   *int    `json:"PriorityID"`
	OwnerID      *int    `json:"OwnerID"`
	LockID       *int    `json:"LockID"`
}

type VoteTicket struct {
	StateID *int `json:"StateID"`
}

type Attachment struct {
	Content     string `json:"Content"`
	ContentType string `json:"ContentType"`
	Filename    string `json:"Filename"`
}

type DynamicField struct {
	Name  string      `json:"Name"`
	Value interface{} `json:"Value"`
}

type OtrsRequest struct {
	TicketID     *string       `json:"TicketID"`
	Ticket       *Ticket       `json:"Ticket"`
	Article      *Article      `json:"Article"`
	Attachment   *Attachment   `json:"Attachment"`
	DynamicField *DynamicField `json:"DynamicField"`
	UserLogin    string
	Password     string
}

type OtrsResponse struct {
	ArticleID    string      `json:"ArticleID"`
	TicketNumber string      `json:"TicketNumber"`
	TicketID     interface{} `json:"TicketID"`
}

type OtrsConfirmRequest struct {
	Email         string `json:"Email"`
	Code          int    `json:"Code"`
	TelegramLogin string `json:"TelegramLogin"`
	UserLogin     string `json:"UserLogin"`
	Password      string `json:"Password"`
}

type OtrsConfirmResponse struct {
	Data      map[string]string `json:"data"`
	Sent      int               `json:"sent"`
	UserLogin string            `json:"user_login"`
}

type OtrsVoteRequest struct {
	TicketID     interface{}  `json:"TicketID"`
	DynamicField DynamicField `json:"DynamicField"`
	Ticket       *VoteTicket  `json:"Ticket"`
}
