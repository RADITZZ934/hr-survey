package models

type SurveyCategory struct {
	BaseModel
	Name        string `gorm:"size:100;uniqueIndex;not null" json:"name"`
	Description string `gorm:"type:text" json:"description"`
}
