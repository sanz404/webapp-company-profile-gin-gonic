package models

import (
	"time"
)

type Product struct {
	Id          uint64    `json:"id" gorm:"primary_key"`
	CategoryId  uint64    `json:"category_id" gorm:"index;not null"`
	Image       string    `json:"image" gorm:"type:text"`
	Name        string    `json:"name" gorm:"index;size:64;not null"`
	Price       float64   `json:"price" gorm:"index;default:0;type:decimal(18,4)"`
	Description string    `json:"description"  gorm:"type:text"`
	IsPublished uint8     `json:"is_published" gorm:"index;default:0"`
	CreatedAt   time.Time `gorm:"index;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"index;default:CURRENT_TIMESTAMP" json:"updated_at"`
}
