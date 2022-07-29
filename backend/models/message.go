package models

import (
	"time"
)

type Message struct {
	Id        uint64    `json:"id" gorm:"primary_key"`
	FullName  string    `json:"full_name" gorm:"index;size:64;not null"`
	Email     string    `json:"email" gorm:"index;size:191;not null"`
	Phone     string    `json:"phone" gorm:"index;size:191;not null"`
	Message   string    `json:"message"  gorm:"type:text"`
	CreatedAt time.Time `gorm:"index;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"index;default:CURRENT_TIMESTAMP" json:"updated_at"`
}
