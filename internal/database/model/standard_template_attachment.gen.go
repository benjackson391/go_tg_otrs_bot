// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameStandardTemplateAttachment = "standard_template_attachment"

// StandardTemplateAttachment mapped from table <standard_template_attachment>
type StandardTemplateAttachment struct {
	ID                   int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	StandardAttachmentID int32     `gorm:"column:standard_attachment_id;not null" json:"standard_attachment_id"`
	StandardTemplateID   int32     `gorm:"column:standard_template_id;not null" json:"standard_template_id"`
	CreateTime           time.Time `gorm:"column:create_time;not null" json:"create_time"`
	CreateBy             int32     `gorm:"column:create_by;not null" json:"create_by"`
	ChangeTime           time.Time `gorm:"column:change_time;not null" json:"change_time"`
	ChangeBy             int32     `gorm:"column:change_by;not null" json:"change_by"`
}

// TableName StandardTemplateAttachment's table name
func (*StandardTemplateAttachment) TableName() string {
	return TableNameStandardTemplateAttachment
}
