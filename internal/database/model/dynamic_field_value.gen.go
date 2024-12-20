// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameDynamicFieldValue = "dynamic_field_value"

// DynamicFieldValue mapped from table <dynamic_field_value>
type DynamicFieldValue struct {
	ID        int32      `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	FieldID   int32      `gorm:"column:field_id;not null" json:"field_id"`
	ObjectID  int64      `gorm:"column:object_id;not null" json:"object_id"`
	ValueText *string    `gorm:"column:value_text" json:"value_text"`
	ValueDate *time.Time `gorm:"column:value_date" json:"value_date"`
	ValueInt  *int64     `gorm:"column:value_int" json:"value_int"`
}

// TableName DynamicFieldValue's table name
func (*DynamicFieldValue) TableName() string {
	return TableNameDynamicFieldValue
}
