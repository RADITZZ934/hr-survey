package models

// SurveyResponse menyimpan data hasil survey dari customer via QR code toko.
// Tabel ini terpisah dari tabel 'responses' yang digunakan untuk survey internal HR.
type SurveyResponse struct {
	BaseModel
	SurveyID      uint   `json:"survey_id" gorm:"not null;index"`
	IDStore       string `json:"id_store" gorm:"size:100;not null"`
	NamaResponden string `json:"nama_responden" gorm:"size:255;not null"`
	Penilaian     string `json:"penilaian" gorm:"size:50;not null"`
}
