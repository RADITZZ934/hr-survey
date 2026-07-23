package models

type Question struct {
	BaseModel
	CategoryID uint           `json:"category_id"`
	Category   SurveyCategory `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
	Text       string         `gorm:"type:text;not null" json:"text"`
	Type       string         `gorm:"size:50;not null;default:'scale'" json:"type"` // scale, text, radio
	IsRequired bool           `gorm:"default:true" json:"is_required"`
	Options    string         `gorm:"type:text" json:"options,omitempty"` // JSON string for multiple choice options if any
}