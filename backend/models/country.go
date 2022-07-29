package models

import (
	"time"
)

type Country struct {
	Id        uint64    `json:"id" gorm:"primary_key"`
	Code      string    `json:"code" gorm:"index;size:64;not null"`
	Name      string    `json:"name" gorm:"index;size:64;not null"`
	CreatedAt time.Time `gorm:"index;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"index;default:CURRENT_TIMESTAMP" json:"updated_at"`
}
