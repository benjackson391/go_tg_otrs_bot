// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameArticleSenderType = "article_sender_type"

// ArticleSenderType mapped from table <article_sender_type>
type ArticleSenderType struct {
	ID         int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name       string    `gorm:"column:name;not null" json:"name"`
	Comments   *string   `gorm:"column:comments" json:"comments"`
	ValidID    int32     `gorm:"column:valid_id;not null" json:"valid_id"`
	CreateTime time.Time `gorm:"column:create_time;not null" json:"create_time"`
	CreateBy   int32     `gorm:"column:create_by;not null" json:"create_by"`
	ChangeTime time.Time `gorm:"column:change_time;not null" json:"change_time"`
	ChangeBy   int32     `gorm:"column:change_by;not null" json:"change_by"`
}

// TableName ArticleSenderType's table name
func (*ArticleSenderType) TableName() string {
	return TableNameArticleSenderType
}