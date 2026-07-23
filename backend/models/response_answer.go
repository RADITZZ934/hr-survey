package models

type ResponseAnswer struct {
	BaseModel
	ResponseID uint     `json:"response_id"`
	Response   Response `gorm:"foreignKey:ResponseID" json:"response,omitempty"`
	QuestionID uint     `json:"question_id"`
	Question   Question `gorm:"foreignKey:QuestionID" json:"question,omitempty"`
	Score      *int     `json:"score,omitempty"` // For scale ratings (e.g., 1-5)
	AnswerText string   `gorm:"type:text" json:"answer_text,omitempty"` // For open-ended questions
}
