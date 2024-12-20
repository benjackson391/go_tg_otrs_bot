// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameFormDraft = "form_draft"

// FormDraft mapped from table <form_draft>
type FormDraft struct {
	ID         int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	ObjectType string    `gorm:"column:object_type;not null" json:"object_type"`
	ObjectID   int32     `gorm:"column:object_id;not null" json:"object_id"`
	Action     string    `gorm:"column:action;not null" json:"action"`
	Title      *string   `gorm:"column:title" json:"title"`
	Content    []byte    `gorm:"column:content;not null" json:"content"`
	CreateTime time.Time `gorm:"column:create_time;not null" json:"create_time"`
	CreateBy   int32     `gorm:"column:create_by;not null" json:"create_by"`
	ChangeTime time.Time `gorm:"column:change_time;not null" json:"change_time"`
	ChangeBy   int32     `gorm:"column:change_by;not null" json:"change_by"`
}

// TableName FormDraft's table name
func (*FormDraft) TableName() string {
	return TableNameFormDraft
}
