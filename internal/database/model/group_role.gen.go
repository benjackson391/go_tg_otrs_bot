// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameGroupRole = "group_role"

// GroupRole mapped from table <group_role>
type GroupRole struct {
	RoleID          int32     `gorm:"column:role_id;not null" json:"role_id"`
	GroupID         int32     `gorm:"column:group_id;not null" json:"group_id"`
	PermissionKey   string    `gorm:"column:permission_key;not null" json:"permission_key"`
	PermissionValue int32     `gorm:"column:permission_value;not null" json:"permission_value"`
	CreateTime      time.Time `gorm:"column:create_time;not null" json:"create_time"`
	CreateBy        int32     `gorm:"column:create_by;not null" json:"create_by"`
	ChangeTime      time.Time `gorm:"column:change_time;not null" json:"change_time"`
	ChangeBy        int32     `gorm:"column:change_by;not null" json:"change_by"`
}

// TableName GroupRole's table name
func (*GroupRole) TableName() string {
	return TableNameGroupRole
}
