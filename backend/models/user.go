package models

import (
	"time"
)

type User struct {
	Id            uint64    `json:"id" gorm:"primary_key"`
	CountryId     uint64    `json:"user_id" gorm:"index;"`
	Username      string    `json:"username" gorm:"index;size:64;not null"`
	Email         string    `json:"email" gorm:"index;size:64;not null"`
	Phone         string    `json:"phone" gorm:"index;size:64;"`
	Address1      string    `json:"address1"  gorm:"type:text"`
	Address2      string    `json:"address2"  gorm:"type:text"`
	City          string    `json:"city" gorm:"index;size:191"`
	ZipCode       string    `json:"zip_code" gorm:"index;size:191"`
	Password      string    `json:"password" gorm:"index;size:191;not null"`
	PasswordToken string    `json:"password_reset_token" gorm:"index;size:191"`
	IsAdmin       uint8     `json:"is_admin" gorm:"index;default:0"`
	Status        uint8     `json:"status" gorm:"index;default:0"`
	CreatedAt     time.Time `gorm:"index;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     time.Time `gorm:"index;default:CURRENT_TIMESTAMP" json:"updated_at"`
}
