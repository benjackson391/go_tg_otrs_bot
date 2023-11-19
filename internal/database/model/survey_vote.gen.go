// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameSurveyVote = "survey_vote"

// SurveyVote mapped from table <survey_vote>
type SurveyVote struct {
	ID         int64      `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	RequestID  int64      `gorm:"column:request_id;not null" json:"request_id"`
	QuestionID int64      `gorm:"column:question_id;not null" json:"question_id"`
	VoteValue  string     `gorm:"column:vote_value;not null" json:"vote_value"`
	CreateTime *time.Time `gorm:"column:create_time" json:"create_time"`
}

// TableName SurveyVote's table name
func (*SurveyVote) TableName() string {
	return TableNameSurveyVote
}
