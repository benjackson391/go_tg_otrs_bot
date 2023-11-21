// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNamePackageRepository = "package_repository"

// PackageRepository mapped from table <package_repository>
type PackageRepository struct {
	ID            int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name          string    `gorm:"column:name;not null" json:"name"`
	Version       string    `gorm:"column:version;not null" json:"version"`
	Vendor        string    `gorm:"column:vendor;not null" json:"vendor"`
	InstallStatus string    `gorm:"column:install_status;not null" json:"install_status"`
	Filename      *string   `gorm:"column:filename" json:"filename"`
	ContentType   *string   `gorm:"column:content_type" json:"content_type"`
	Content       []byte    `gorm:"column:content;not null" json:"content"`
	CreateTime    time.Time `gorm:"column:create_time;not null" json:"create_time"`
	CreateBy      int32     `gorm:"column:create_by;not null" json:"create_by"`
	ChangeTime    time.Time `gorm:"column:change_time;not null" json:"change_time"`
	ChangeBy      int32     `gorm:"column:change_by;not null" json:"change_by"`
}

// TableName PackageRepository's table name
func (*PackageRepository) TableName() string {
	return TableNamePackageRepository
}