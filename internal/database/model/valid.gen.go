// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameValid = "valid"

// Valid mapped from table <valid>
type Valid struct {
	ID         int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name       string    `gorm:"column:name;not null" json:"name"`
	CreateTime time.Time `gorm:"column:create_time;not null" json:"create_time"`
	CreateBy   int32     `gorm:"column:create_by;not null" json:"create_by"`
	ChangeTime time.Time `gorm:"column:change_time;not null" json:"change_time"`
	ChangeBy   int32     `gorm:"column:change_by;not null" json:"change_by"`
}

// TableName Valid's table name
func (*Valid) TableName() string {
	return TableNameValid
}
