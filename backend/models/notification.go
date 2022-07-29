package models

import (
	"time"
)

type Notification struct {
	Id          uint64    `json:"id" gorm:"primary_key"`
	UserId      uint64    `json:"user_id" gorm:"index;not null"`
	Subject     string    `json:"subject" gorm:"index;size:191;not null"`
	SortContent string    `json:"sort_content" gorm:"index;size:191;not null"`
	FullContent string    `json:"full_content"  gorm:"type:longtext"`
	ReadedAt    time.Time `json:"ReadedAt" gorm:"index"`
	CreatedAt   time.Time `gorm:"index;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"index;default:CURRENT_TIMESTAMP" json:"updated_at"`
}
