// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameAclSync = "acl_sync"

// AclSync mapped from table <acl_sync>
type AclSync struct {
	AclID      string    `gorm:"column:acl_id;not null" json:"acl_id"`
	SyncState  string    `gorm:"column:sync_state;not null" json:"sync_state"`
	CreateTime time.Time `gorm:"column:create_time;not null" json:"create_time"`
	ChangeTime time.Time `gorm:"column:change_time;not null" json:"change_time"`
}

// TableName AclSync's table name
func (*AclSync) TableName() string {
	return TableNameAclSync
}
