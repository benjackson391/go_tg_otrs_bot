// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameDynamicFieldObjIDName = "dynamic_field_obj_id_name"

// DynamicFieldObjIDName mapped from table <dynamic_field_obj_id_name>
type DynamicFieldObjIDName struct {
	ObjectID   int32  `gorm:"column:object_id;primaryKey;autoIncrement:true" json:"object_id"`
	ObjectName string `gorm:"column:object_name;not null" json:"object_name"`
	ObjectType string `gorm:"column:object_type;not null" json:"object_type"`
}

// TableName DynamicFieldObjIDName's table name
func (*DynamicFieldObjIDName) TableName() string {
	return TableNameDynamicFieldObjIDName
}
