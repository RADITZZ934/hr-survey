package models

type SurveyTemplate struct {
	BaseModel
	Title       string     `gorm:"size:150;not null" json:"title"`
	Description string     `gorm:"type:text" json:"description"`
	CreatedBy   uint       `json:"created_by"`
	User        User       `gorm:"foreignKey:CreatedBy" json:"creator,omitempty"`
}

func (SurveyTemplate) TableName() string {
	return "survey_templates"
}
