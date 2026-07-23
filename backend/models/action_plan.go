package models

import (
	"time"
)

type ActionPlan struct {
	BaseModel
	SurveyID    uint       `json:"survey_id"`
	Survey      Survey     `gorm:"foreignKey:SurveyID" json:"survey,omitempty"`
	Title       string     `gorm:"size:150;not null" json:"title"`
	Description string     `gorm:"type:text" json:"description"`
	Status      string     `gorm:"size:50;not null;default:'pending'" json:"status"` // pending, in_progress, completed
	TargetDate  *time.Time `json:"target_date,omitempty"`
	AssigneeID  *uint      `json:"assignee_id,omitempty"`
	Assignee    *User      `gorm:"foreignKey:AssigneeID" json:"assignee,omitempty"`
	CreatedBy   uint       `json:"created_by"`
	Creator     User       `gorm:"foreignKey:CreatedBy" json:"creator,omitempty"`
}
