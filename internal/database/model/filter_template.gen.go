// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameFilterTemplate = "filter_template"

// FilterTemplate mapped from table <filter_template>
type FilterTemplate struct {
	ID     int32   `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name   string  `gorm:"column:name;not null" json:"name"`
	Filter *string `gorm:"column:filter" json:"filter"`
	StatID *int32  `gorm:"column:stat_id" json:"stat_id"`
}

// TableName FilterTemplate's table name
func (*FilterTemplate) TableName() string {
	return TableNameFilterTemplate
}
