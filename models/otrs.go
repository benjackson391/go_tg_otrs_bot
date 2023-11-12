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
	Title        *string `json:"Title,omitempty"`
	QueueID      *int    `json:"QueueID,omitempty"`
	TypeID       *int    `json:"TypeID,omitempty"`
	CustomerUser *string `json:"CustomerUser,omitempty"`
	StateID      *int    `json:"StateID,omitempty"`
	PriorityID   *int    `json:"PriorityID,omitempty"`
	OwnerID      *int    `json:"OwnerID,omitempty"`
	LockID       *int    `json:"LockID,omitempty"`
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
	TicketID     *string       `json:"TicketID,omitempty"`
	Ticket       *Ticket       `json:"Ticket,omitempty"`
	Article      *Article      `json:"Article,omitempty"`
	Attachment   *Attachment   `json:"Attachment,omitempty"`
	DynamicField *DynamicField `json:"DynamicField,omitempty"`
	UserLogin    string        `json:"UserLogin"`
	Password     string        `json:"Password"`
}

type OtrsResponse struct {
	ArticleID    string      `json:"ArticleID"`
	TicketNumber string      `json:"TicketNumber"`
	TicketID     interface{} `json:"TicketID"`
}

type OtrsConfirmRequest struct {
	Email         string `json:"Email,omitempty"`
	Code          int    `json:"Code,omitempty"`
	TelegramLogin string `json:"TelegramLogin,omitempty"`
	UserLogin     string `json:"UserLogin"`
	Password      string `json:"Password"`
}

type OtrsConfirmResponse struct {
	Data      map[string]string `json:"data"`
	Sent      int               `json:"sent"`
	UserLogin string            `json:"user_login"`
}
