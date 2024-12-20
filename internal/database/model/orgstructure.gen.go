// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameOrgstructure = "orgstructure"

// Orgstructure mapped from table <orgstructure>
type Orgstructure struct {
	OrgstructureID int32   `gorm:"column:orgstructure_id;primaryKey;autoIncrement:true" json:"orgstructure_id"`
	Name           string  `gorm:"column:name;not null" json:"name"`
	Comments       *string `gorm:"column:comments" json:"comments"`
	ParentID       *int32  `gorm:"column:parent_id" json:"parent_id"`
	OwnerID        *int32  `gorm:"column:owner_id" json:"owner_id"`
	Members        *string `gorm:"column:members" json:"members"`
}

// TableName Orgstructure's table name
func (*Orgstructure) TableName() string {
	return TableNameOrgstructure
}
