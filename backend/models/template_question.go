package models

type TemplateQuestion struct {
	BaseModel
	TemplateID uint           `gorm:"index" json:"template_id"`
	Template   SurveyTemplate `gorm:"foreignKey:TemplateID" json:"template,omitempty"`
	QuestionID uint           `gorm:"index" json:"question_id"`
	Question   Question       `gorm:"foreignKey:QuestionID" json:"question,omitempty"`
}
