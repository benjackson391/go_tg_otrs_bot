// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameSurveyAnswer = "survey_answer"

// SurveyAnswer mapped from table <survey_answer>
type SurveyAnswer struct {
	ID         int64      `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	QuestionID int64      `gorm:"column:question_id;not null" json:"question_id"`
	Answer     string     `gorm:"column:answer;not null" json:"answer"`
	Position   int32      `gorm:"column:position;not null" json:"position"`
	CreateTime *time.Time `gorm:"column:create_time" json:"create_time"`
	CreateBy   *int32     `gorm:"column:create_by" json:"create_by"`
	ChangeTime *time.Time `gorm:"column:change_time" json:"change_time"`
	ChangeBy   *int32     `gorm:"column:change_by" json:"change_by"`
}

// TableName SurveyAnswer's table name
func (*SurveyAnswer) TableName() string {
	return TableNameSurveyAnswer
}
