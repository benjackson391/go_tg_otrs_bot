// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTicketFlag = "ticket_flag"

// TicketFlag mapped from table <ticket_flag>
type TicketFlag struct {
	TicketID    int64     `gorm:"column:ticket_id;primaryKey" json:"ticket_id"`
	TicketKey   string    `gorm:"column:ticket_key;primaryKey" json:"ticket_key"`
	TicketValue *string   `gorm:"column:ticket_value" json:"ticket_value"`
	CreateTime  time.Time `gorm:"column:create_time;not null" json:"create_time"`
	CreateBy    int32     `gorm:"column:create_by;primaryKey" json:"create_by"`
}

// TableName TicketFlag's table name
func (*TicketFlag) TableName() string {
	return TableNameTicketFlag
}
