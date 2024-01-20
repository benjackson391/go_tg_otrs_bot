package models

type StateMachine struct {
	State   string    `json:"state"`
	Type    string    `json:"type"`
	Waiting *[]string `json:"waiting"`
}
