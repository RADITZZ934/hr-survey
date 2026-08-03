package models

import (
	"time"
)

type Survey struct {
	BaseModel
	Visibility  string         `gorm:"size:20;not null;default:'internal'" json:"visibility"` // internal, external
	TemplateID  *uint          `json:"template_id" gorm:"default:null"`
	Template    SurveyTemplate `gorm:"foreignKey:TemplateID" json:"template,omitempty"`
	Title       string         `gorm:"size:150;not null" json:"title"`
	Description string         `gorm:"type:text" json:"description"`
	StartDate   time.Time      `json:"start_date"`
	EndDate     time.Time      `json:"end_date"`
	Status      string         `gorm:"size:50;not null;default:'draft'" json:"status"` // draft, active, closed
	CreatedBy      uint           `json:"created_by"`
	Creator        User           `gorm:"foreignKey:CreatedBy" json:"creator,omitempty"`
	ResponsesCount int64          `json:"responses_count" gorm:"-"`
}
