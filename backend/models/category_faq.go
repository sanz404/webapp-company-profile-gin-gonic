package models

import (
	"time"
)

type CategoryFaq struct {
	Id          uint64    `json:"id" gorm:"primary_key"`
	Name        string    `json:"name" gorm:"index;size:64;not null"`
	Description string    `json:"description"  gorm:"type:text"`
	CreatedAt   time.Time `gorm:"index;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"index;default:CURRENT_TIMESTAMP" json:"updated_at"`
}
