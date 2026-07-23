package models

type ActivityLog struct {
	BaseModel
	UserID      *uint  `json:"user_id,omitempty"`
	User        *User  `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Action      string `gorm:"size:100;not null" json:"action"`
	Description string `gorm:"type:text" json:"description"`
	IPAddress   string `gorm:"size:45" json:"ip_address,omitempty"`
}
