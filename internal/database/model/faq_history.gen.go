// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameFaqHistory = "faq_history"

// FaqHistory mapped from table <faq_history>
type FaqHistory struct {
	ID        int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name      string    `gorm:"column:name;not null" json:"name"`
	ItemID    int32     `gorm:"column:item_id;not null" json:"item_id"`
	Created   time.Time `gorm:"column:created;not null" json:"created"`
	CreatedBy int32     `gorm:"column:created_by;not null" json:"created_by"`
	Changed   time.Time `gorm:"column:changed;not null" json:"changed"`
	ChangedBy int32     `gorm:"column:changed_by;not null" json:"changed_by"`
}

// TableName FaqHistory's table name
func (*FaqHistory) TableName() string {
	return TableNameFaqHistory
}
