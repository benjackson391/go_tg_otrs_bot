// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameFaqItem = "faq_item"

// FaqItem mapped from table <faq_item>
type FaqItem struct {
	ID          int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	FNumber     string    `gorm:"column:f_number;not null" json:"f_number"`
	FSubject    *string   `gorm:"column:f_subject" json:"f_subject"`
	FName       string    `gorm:"column:f_name;not null" json:"f_name"`
	FLanguageID int32     `gorm:"column:f_language_id;not null" json:"f_language_id"`
	StateID     int32     `gorm:"column:state_id;not null" json:"state_id"`
	CategoryID  int32     `gorm:"column:category_id;not null" json:"category_id"`
	Approved    int32     `gorm:"column:approved;not null;default:1" json:"approved"`
	ValidID     int32     `gorm:"column:valid_id;not null;default:1" json:"valid_id"`
	ContentType string    `gorm:"column:content_type;not null;default:text/html" json:"content_type"`
	FKeywords   *string   `gorm:"column:f_keywords" json:"f_keywords"`
	FField1     *string   `gorm:"column:f_field1" json:"f_field1"`
	FField2     *string   `gorm:"column:f_field2" json:"f_field2"`
	FField3     *string   `gorm:"column:f_field3" json:"f_field3"`
	FField4     *string   `gorm:"column:f_field4" json:"f_field4"`
	FField5     *string   `gorm:"column:f_field5" json:"f_field5"`
	FField6     *string   `gorm:"column:f_field6" json:"f_field6"`
	Created     time.Time `gorm:"column:created;not null" json:"created"`
	CreatedBy   int32     `gorm:"column:created_by;not null" json:"created_by"`
	Changed     time.Time `gorm:"column:changed;not null" json:"changed"`
	ChangedBy   int32     `gorm:"column:changed_by;not null" json:"changed_by"`
}

// TableName FaqItem's table name
func (*FaqItem) TableName() string {
	return TableNameFaqItem
}
