package models

type SurveyQuestion struct {
	BaseModel
	SurveyID   uint     `gorm:"index" json:"survey_id"`
	Survey     Survey   `gorm:"foreignKey:SurveyID" json:"survey,omitempty"`
	QuestionID uint     `gorm:"index" json:"question_id"`
	Question   Question `gorm:"foreignKey:QuestionID" json:"question,omitempty"`
}
