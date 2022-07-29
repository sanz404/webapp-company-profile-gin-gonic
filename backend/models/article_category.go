package models

type ArticleCategory struct {
	Id         uint64 `json:"id" gorm:"primary_key"`
	ArticleId  uint64 `json:"article_id" gorm:"index;not null"`
	CategoryId uint64 `json:"category_id" gorm:"index;not null"`
}
