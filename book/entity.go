package book

import "gorm.io/gorm"

type Books struct {
	gorm.Model
	Title     string `json:title`
	Year      int    `json:year`
	Publisher string `json:publisher`
	UserID    uint   `json:"user_id"`
}
