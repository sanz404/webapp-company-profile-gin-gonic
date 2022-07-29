package models

type ArticleCategory struct {
	Id        uint64 `json:"id" gorm:"primary_key"`
	UserId    uint64 `json:"user_id" gorm:"index;not null"`
	ArticleId uint64 `json:"Article_id" gorm:"index;not null"`
}
