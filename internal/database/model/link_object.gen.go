// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameLinkObject = "link_object"

// LinkObject mapped from table <link_object>
type LinkObject struct {
	ID   int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name string `gorm:"column:name;not null" json:"name"`
}

// TableName LinkObject's table name
func (*LinkObject) TableName() string {
	return TableNameLinkObject
}