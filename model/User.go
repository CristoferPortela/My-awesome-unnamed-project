package model

import "time"

type User struct {
	ID           uint   `gorm:"primaryKey"`
	Name         string `gorm:"not null"`
	Email        string `gorm:"not null"`
	Password     string `gorm:"not null"`
	Username     string `gorm:"not null"`
	IsAdmin      bool   `gorm:"type:bool;default:0"`
	IsSuperAdmin bool   `gorm:"type:bool;default:0"`
	IsActivated  bool   `gorm:"type:bool;default:0"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
