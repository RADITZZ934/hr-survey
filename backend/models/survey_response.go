package models

// SurveyResponse menyimpan data hasil survey dari customer via QR code toko.
// Tabel ini terpisah dari tabel 'responses' yang digunakan untuk survey internal HR.
type SurveyResponse struct {
	BaseModel
	SurveyID       uint   `json:"survey_id" gorm:"not null;index"`
	IDStore        string `json:"id_store" gorm:"size:100;not null"`
	NamaResponden  string `json:"nama_responden" gorm:"size:255;not null"`
	EmailResponden string `json:"email_responden" gorm:"size:150"`
	Penilaian      string `json:"penilaian" gorm:"size:50;not null"`
	IPAddress      string `json:"ip_address" gorm:"size:50"`
	IsFlagged      bool   `json:"is_flagged" gorm:"default:false"`
	FlagReason     string `json:"flag_reason" gorm:"size:150"`
}
