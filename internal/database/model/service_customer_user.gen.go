// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameServiceCustomerUser = "service_customer_user"

// ServiceCustomerUser mapped from table <service_customer_user>
type ServiceCustomerUser struct {
	CustomerUserLogin string    `gorm:"column:customer_user_login;not null" json:"customer_user_login"`
	ServiceID         int32     `gorm:"column:service_id;not null" json:"service_id"`
	CreateTime        time.Time `gorm:"column:create_time;not null" json:"create_time"`
	CreateBy          int32     `gorm:"column:create_by;not null" json:"create_by"`
}

// TableName ServiceCustomerUser's table name
func (*ServiceCustomerUser) TableName() string {
	return TableNameServiceCustomerUser
}