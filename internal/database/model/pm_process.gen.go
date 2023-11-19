// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNamePmProcess = "pm_process"

// PmProcess mapped from table <pm_process>
type PmProcess struct {
	ID            int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	EntityID      string    `gorm:"column:entity_id;not null" json:"entity_id"`
	Name          string    `gorm:"column:name;not null" json:"name"`
	StateEntityID string    `gorm:"column:state_entity_id;not null" json:"state_entity_id"`
	Layout        []byte    `gorm:"column:layout;not null" json:"layout"`
	Config        []byte    `gorm:"column:config;not null" json:"config"`
	CreateTime    time.Time `gorm:"column:create_time;not null" json:"create_time"`
	CreateBy      int32     `gorm:"column:create_by;not null" json:"create_by"`
	ChangeTime    time.Time `gorm:"column:change_time;not null" json:"change_time"`
	ChangeBy      int32     `gorm:"column:change_by;not null" json:"change_by"`
}

// TableName PmProcess's table name
func (*PmProcess) TableName() string {
	return TableNamePmProcess
}
