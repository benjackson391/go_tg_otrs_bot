// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNamePostmasterFilter = "postmaster_filter"

// PostmasterFilter mapped from table <postmaster_filter>
type PostmasterFilter struct {
	FName  string `gorm:"column:f_name;not null" json:"f_name"`
	FStop  *int32 `gorm:"column:f_stop" json:"f_stop"`
	FType  string `gorm:"column:f_type;not null" json:"f_type"`
	FKey   string `gorm:"column:f_key;not null" json:"f_key"`
	FValue string `gorm:"column:f_value;not null" json:"f_value"`
	FNot   *int32 `gorm:"column:f_not" json:"f_not"`
}

// TableName PostmasterFilter's table name
func (*PostmasterFilter) TableName() string {
	return TableNamePostmasterFilter
}