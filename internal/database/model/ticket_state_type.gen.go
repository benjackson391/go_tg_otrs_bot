// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTicketStateType = "ticket_state_type"

// TicketStateType mapped from table <ticket_state_type>
type TicketStateType struct {
	ID         int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name       string    `gorm:"column:name;not null" json:"name"`
	Comments   *string   `gorm:"column:comments" json:"comments"`
	CreateTime time.Time `gorm:"column:create_time;not null" json:"create_time"`
	CreateBy   int32     `gorm:"column:create_by;not null" json:"create_by"`
	ChangeTime time.Time `gorm:"column:change_time;not null" json:"change_time"`
	ChangeBy   int32     `gorm:"column:change_by;not null" json:"change_by"`
}

// TableName TicketStateType's table name
func (*TicketStateType) TableName() string {
	return TableNameTicketStateType
}
