// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameVirtualFsPreference = "virtual_fs_preferences"

// VirtualFsPreference mapped from table <virtual_fs_preferences>
type VirtualFsPreference struct {
	VirtualFsID      int64   `gorm:"column:virtual_fs_id;not null" json:"virtual_fs_id"`
	PreferencesKey   string  `gorm:"column:preferences_key;not null" json:"preferences_key"`
	PreferencesValue *string `gorm:"column:preferences_value" json:"preferences_value"`
}

// TableName VirtualFsPreference's table name
func (*VirtualFsPreference) TableName() string {
	return TableNameVirtualFsPreference
}