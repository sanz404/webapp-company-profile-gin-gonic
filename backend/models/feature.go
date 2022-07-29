package models

import (
	"time"
)

type Feature struct {
	Id          uint64    `json:"id" gorm:"primary_key"`
	Icon        string    `json:"icon" gorm:"type:text"`
	Title       string    `json:"title" gorm:"index;size:191;not null"`
	Description string    `json:"description"  gorm:"type:text"`
	IsPublished uint8     `json:"is_published" gorm:"index;default:0"`
	CreatedAt   time.Time `gorm:"index;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"index;default:CURRENT_TIMESTAMP" json:"updated_at"`
}
