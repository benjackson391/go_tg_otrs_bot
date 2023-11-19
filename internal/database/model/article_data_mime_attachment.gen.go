// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameArticleDataMimeAttachment = "article_data_mime_attachment"

// ArticleDataMimeAttachment mapped from table <article_data_mime_attachment>
type ArticleDataMimeAttachment struct {
	ID                 int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	ArticleID          int64     `gorm:"column:article_id;not null" json:"article_id"`
	Filename           *string   `gorm:"column:filename" json:"filename"`
	ContentSize        *string   `gorm:"column:content_size" json:"content_size"`
	ContentType        *string   `gorm:"column:content_type" json:"content_type"`
	ContentID          *string   `gorm:"column:content_id" json:"content_id"`
	ContentAlternative *string   `gorm:"column:content_alternative" json:"content_alternative"`
	Disposition        *string   `gorm:"column:disposition" json:"disposition"`
	Content            *[]byte   `gorm:"column:content" json:"content"`
	CreateTime         time.Time `gorm:"column:create_time;not null" json:"create_time"`
	CreateBy           int32     `gorm:"column:create_by;not null" json:"create_by"`
	ChangeTime         time.Time `gorm:"column:change_time;not null" json:"change_time"`
	ChangeBy           int32     `gorm:"column:change_by;not null" json:"change_by"`
}

// TableName ArticleDataMimeAttachment's table name
func (*ArticleDataMimeAttachment) TableName() string {
	return TableNameArticleDataMimeAttachment
}
