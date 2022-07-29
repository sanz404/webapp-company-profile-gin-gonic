package models

import (
	"time"
)

type Faq struct {
	Id          uint64    `json:"id" gorm:"primary_key"`
	CategoryId  uint64    `json:"category_id" gorm:"index;not null"`
	Question    string    `json:"name" gorm:"index;size:191;not null"`
	Answer      string    `json:"answer"  gorm:"type:text"`
	Sort        uint      `json:"sort" gorm:"index;default:0"`
	IsPublished uint8     `json:"is_published" gorm:"index;default:0"`
	CreatedAt   time.Time `gorm:"index;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"index;default:CURRENT_TIMESTAMP" json:"updated_at"`
}
