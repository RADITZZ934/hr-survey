package models

type User struct {
	BaseModel
	Username   string   `gorm:"size:100;uniqueIndex;not null" json:"username"`
	Email      string   `gorm:"size:100;uniqueIndex;not null" json:"email"`
	Password   string   `gorm:"size:255;not null" json:"-"`
	Role       string   `gorm:"size:50;not null;default:'employee'" json:"role"` // admin, hr, manager, employee
	Department string   `gorm:"size:100;default:''" json:"department"`
	AvgScore   *float64 `gorm:"-" json:"avg_score"`
	Percentage *int     `gorm:"-" json:"percentage"`
}
