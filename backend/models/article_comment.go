package models

import (
	"time"
)

type ArticleComment struct {
	Id          uint64    `json:"id" gorm:"primary_key"`
	UserId      uint64    `json:"user_id" gorm:"index;not null"`
	ArticleId   uint64    `json:"article_id" gorm:"index;not null"`
	Description string    `json:"description"  gorm:"type:text"`
	CreatedAt   time.Time `gorm:"index;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"index;default:CURRENT_TIMESTAMP" json:"updated_at"`
}
