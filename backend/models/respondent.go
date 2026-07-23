package models

import (
	"time"
)

type Respondent struct {
	BaseModel
	SurveyID    uint       `gorm:"index" json:"survey_id"`
	Survey      Survey     `gorm:"foreignKey:SurveyID" json:"survey,omitempty"`
	UserID      uint       `gorm:"index" json:"user_id"`
	User        User       `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Status      string     `gorm:"size:50;not null;default:'pending'" json:"status"` // pending, completed
	CompletedAt *time.Time `json:"completed_at,omitempty"`
}
