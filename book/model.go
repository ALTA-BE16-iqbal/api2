package book

import (
	"log"

	"gorm.io/gorm"
)

type BookModel struct {
	db *gorm.DB
}

func (bm *BookModel) SetDB(db *gorm.DB) {
	bm.db = db
}

func (bm *BookModel) Insert(newBooks Books) (Books, error) {
	if err := bm.db.Create(&newBooks).Error; err != nil {
		log.Panicln("Terjadi error saata create book", err.Error())
	}
	return newBooks, nil
}

func (bm *BookModel) GetAllBook(userID uint) (any, error) {
	type ExpectedRespon struct {
		Title string `json:"title"`
		Year  string `json:"year"`
		Name  string `json:"name"`
	}

	res := []ExpectedRespon{}
	var err error

	if userID != 0 {
		err = bm.db.Table("books").Select("books.title as title, books.year, users.name as name").Joins("JOIN users on users.id = books.user_id").Where("books.user_id = ?", userID).Scan(&res).Error
	} else {
		err = bm.db.Table("books").Select("books.title as title, books.year, users.name as name").Joins("JOIN users on users.id = books.user_id").Scan(&res).Error
	}

	if err != nil {
		log.Println("Terjadi error saat select Book ", err.Error())
		return nil, err
	}
	return res, nil
}

func (bm *BookModel) GetBookByID(id uint) (Books, error) {
	res := Books{}

	if err := bm.db.Find(&res, id).Error; err != nil {
		log.Println("Terjadi error saat select Books ", err.Error())
		return Books{}, err
	}

	return res, nil
}
func (bm *BookModel) DeleteBookByID(id uint) (Books, error) {
	res := Books{}

	if err := bm.db.Find(&res, id).Error; err != nil {
		log.Println("Terjadi error saat select Books ", err.Error())
		return Books{}, err
	}

	if err := bm.db.Delete(&res).Error; err != nil {
		log.Println("Failed to delete", err.Error())
		return Books{}, err
	}

	return res, nil
}
