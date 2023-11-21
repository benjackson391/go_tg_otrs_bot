// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameDynamicField = "dynamic_field"

// DynamicField mapped from table <dynamic_field>
type DynamicField struct {
	ID            int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	InternalField int32     `gorm:"column:internal_field;not null" json:"internal_field"`
	Name          string    `gorm:"column:name;not null" json:"name"`
	Label         string    `gorm:"column:label;not null" json:"label"`
	FieldOrder    int32     `gorm:"column:field_order;not null" json:"field_order"`
	FieldType     string    `gorm:"column:field_type;not null" json:"field_type"`
	ObjectType    string    `gorm:"column:object_type;not null" json:"object_type"`
	Config        *[]byte   `gorm:"column:config" json:"config"`
	ValidID       int32     `gorm:"column:valid_id;not null" json:"valid_id"`
	CreateTime    time.Time `gorm:"column:create_time;not null" json:"create_time"`
	CreateBy      int32     `gorm:"column:create_by;not null" json:"create_by"`
	ChangeTime    time.Time `gorm:"column:change_time;not null" json:"change_time"`
	ChangeBy      int32     `gorm:"column:change_by;not null" json:"change_by"`
}

// TableName DynamicField's table name
func (*DynamicField) TableName() string {
	return TableNameDynamicField
}