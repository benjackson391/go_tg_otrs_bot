// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameCustomerPreference = "customer_preferences"

// CustomerPreference mapped from table <customer_preferences>
type CustomerPreference struct {
	UserID           string  `gorm:"column:user_id;not null" json:"user_id"`
	PreferencesKey   string  `gorm:"column:preferences_key;not null" json:"preferences_key"`
	PreferencesValue *string `gorm:"column:preferences_value" json:"preferences_value"`
}

// TableName CustomerPreference's table name
func (*CustomerPreference) TableName() string {
	return TableNameCustomerPreference
}
