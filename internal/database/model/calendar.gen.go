// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameCalendar = "calendar"

// Calendar mapped from table <calendar>
type Calendar struct {
	ID                 int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	GroupID            int32     `gorm:"column:group_id;not null" json:"group_id"`
	Name               string    `gorm:"column:name;not null" json:"name"`
	SaltString         string    `gorm:"column:salt_string;not null" json:"salt_string"`
	Color              string    `gorm:"column:color;not null" json:"color"`
	TicketAppointments *[]byte   `gorm:"column:ticket_appointments" json:"ticket_appointments"`
	ValidID            int32     `gorm:"column:valid_id;not null" json:"valid_id"`
	CreateTime         time.Time `gorm:"column:create_time;not null" json:"create_time"`
	CreateBy           int32     `gorm:"column:create_by;not null" json:"create_by"`
	ChangeTime         time.Time `gorm:"column:change_time;not null" json:"change_time"`
	ChangeBy           int32     `gorm:"column:change_by;not null" json:"change_by"`
}

// TableName Calendar's table name
func (*Calendar) TableName() string {
	return TableNameCalendar
}