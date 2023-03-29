package user

import (
	"tugas-1-API/book"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	HP       string `json:"hp" gorm:"type:varchar(13);primaryKey"`
	Password string `json:"password"`
	Book     []book.Books
}
