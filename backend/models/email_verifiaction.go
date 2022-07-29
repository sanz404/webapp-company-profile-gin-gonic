package models

import (
	"time"
)

type EmailVerifiaction struct {
	Id             uint64    `json:"id" gorm:"primary_key"`
	UserId         uint64    `json:"user_id" gorm:"index;not null"`
	Code           string    `json:"code" gorm:"index;size:191;not null"`
	Token          string    `json:"token" gorm:"index;size:191;not null"`
	EmailConfirmed uint8     `json:"email_confirmed" gorm:"index;default:0"`
	IsExpired      uint8     `json:"is_expired" gorm:"index;default:0"`
	ExpiredAt      time.Time `json:"expired_at" gorm:"index;not null"`
	CreatedAt      time.Time `gorm:"index;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt      time.Time `gorm:"index;default:CURRENT_TIMESTAMP" json:"updated_at"`
}
