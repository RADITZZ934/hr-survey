package models

import "time"

type Store struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	AplikasiID int       `json:"aplikasi_id" gorm:"default:1"`
	Kode       string    `json:"kode" gorm:"size:50;uniqueIndex"`
	Name       string    `json:"name" gorm:"size:255"`
	Koordinat  string    `json:"koordinat" gorm:"size:100;default:''"`
}
