// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameQueue = "queue"

// Queue mapped from table <queue>
type Queue struct {
	ID                  int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name                string    `gorm:"column:name;not null" json:"name"`
	GroupID             int32     `gorm:"column:group_id;not null" json:"group_id"`
	UnlockTimeout       *int32    `gorm:"column:unlock_timeout" json:"unlock_timeout"`
	FirstResponseTime   *int32    `gorm:"column:first_response_time" json:"first_response_time"`
	FirstResponseNotify *int32    `gorm:"column:first_response_notify" json:"first_response_notify"`
	UpdateTime          *int32    `gorm:"column:update_time" json:"update_time"`
	UpdateNotify        *int32    `gorm:"column:update_notify" json:"update_notify"`
	SolutionTime        *int32    `gorm:"column:solution_time" json:"solution_time"`
	SolutionNotify      *int32    `gorm:"column:solution_notify" json:"solution_notify"`
	SystemAddressID     int32     `gorm:"column:system_address_id;not null" json:"system_address_id"`
	CalendarName        *string   `gorm:"column:calendar_name" json:"calendar_name"`
	DefaultSignKey      *string   `gorm:"column:default_sign_key" json:"default_sign_key"`
	SalutationID        int32     `gorm:"column:salutation_id;not null" json:"salutation_id"`
	SignatureID         int32     `gorm:"column:signature_id;not null" json:"signature_id"`
	FollowUpID          int32     `gorm:"column:follow_up_id;not null" json:"follow_up_id"`
	FollowUpLock        int32     `gorm:"column:follow_up_lock;not null" json:"follow_up_lock"`
	Comments            *string   `gorm:"column:comments" json:"comments"`
	ValidID             int32     `gorm:"column:valid_id;not null" json:"valid_id"`
	CreateTime          time.Time `gorm:"column:create_time;not null" json:"create_time"`
	CreateBy            int32     `gorm:"column:create_by;not null" json:"create_by"`
	ChangeTime          time.Time `gorm:"column:change_time;not null" json:"change_time"`
	ChangeBy            int32     `gorm:"column:change_by;not null" json:"change_by"`
}

// TableName Queue's table name
func (*Queue) TableName() string {
	return TableNameQueue
}
