package models

import (
	"time"
)

type Response struct {
	BaseModel
	SurveyID    uint      `json:"survey_id"`
	Survey      Survey    `gorm:"foreignKey:SurveyID" json:"survey,omitempty"`
	UserID         *uint     `json:"user_id,omitempty"` // Nullable for anonymous responses
	User           *User     `gorm:"foreignKey:UserID" json:"user,omitempty"`
	RespondentID   string    `json:"respondent_id" gorm:"size:100;default:'ANONYMOUS'"`
	RespondentDept string    `json:"respondent_dept" gorm:"size:100;default:'ANONYMOUS'"`
	SubmittedAt    time.Time `json:"submitted_at"`
}
