// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameUserPreference = "user_preferences"

// UserPreference mapped from table <user_preferences>
type UserPreference struct {
	UserID           int32   `gorm:"column:user_id;not null" json:"user_id"`
	PreferencesKey   string  `gorm:"column:preferences_key;not null" json:"preferences_key"`
	PreferencesValue *[]byte `gorm:"column:preferences_value" json:"preferences_value"`
}

// TableName UserPreference's table name
func (*UserPreference) TableName() string {
	return TableNameUserPreference
}
