// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameSession = "sessions"

// Session mapped from table <sessions>
type Session struct {
	ID         int64   `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	SessionID  string  `gorm:"column:session_id;not null" json:"session_id"`
	DataKey    string  `gorm:"column:data_key;not null" json:"data_key"`
	DataValue  *string `gorm:"column:data_value" json:"data_value"`
	Serialized int32   `gorm:"column:serialized;not null" json:"serialized"`
}

// TableName Session's table name
func (*Session) TableName() string {
	return TableNameSession
}