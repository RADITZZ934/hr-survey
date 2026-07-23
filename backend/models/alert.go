package models

import "time"

type Alert struct {
	BaseModel
	SurveyID    uint      `json:"survey_id"`
	Survey      Survey    `gorm:"foreignKey:SurveyID" json:"survey,omitempty"`
	Respondent  string    `gorm:"size:100" json:"respondent"`
	Score       float64   `json:"score"`
	Message     string    `gorm:"type:text" json:"message"`
	IsRead      bool      `gorm:"default:false" json:"is_read"`
	CreatedAt   time.Time `json:"created_at"`
}
