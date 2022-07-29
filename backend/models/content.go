package models

import (
	"time"
)

type Content struct {
	Id        uint64    `json:"id" gorm:"primary_key"`
	KeyName   string    `json:"key_name" gorm:"index;size:191;not null"`
	KeyValue  string    `json:"key_value"  gorm:"type:longtext"`
	CreatedAt time.Time `gorm:"index;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"index;default:CURRENT_TIMESTAMP" json:"updated_at"`
}
