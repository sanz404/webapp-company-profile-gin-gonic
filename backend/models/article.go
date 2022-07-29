package models

import (
	"time"
)

type Article struct {
	Id          uint64    `json:"id" gorm:"primary_key"`
	UserId      uint64    `json:"user_id" gorm:"index;not null"`
	Image       string    `json:"image" gorm:"type:text"`
	Slug        string    `json:"slug" gorm:"index;size:191;not null"`
	Title       string    `json:"title" gorm:"index;size:191;not null"`
	Content     string    `json:"content"  gorm:"type:longtext"`
	IsPublished uint8     `json:"is_published" gorm:"index;default:0"`
	CreatedAt   time.Time `gorm:"index;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"index;default:CURRENT_TIMESTAMP" json:"updated_at"`
}
