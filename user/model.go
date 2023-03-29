package user

import (
	"errors"
	"log"

	"gorm.io/gorm"
)

type UserModel struct {
	db *gorm.DB
}

func (um *UserModel) SetDB(db *gorm.DB) {
	um.db = db
}

func (um *UserModel) Insert(newUser User) (User, error) {
	if err := um.db.Create(&newUser).Error; err != nil {
		log.Println("Terjadi error saat create user", err.Error())
		return User{}, err
	}

	return newUser, nil
}

func (um *UserModel) Login(hp, password string) (User, error) {
	res := User{}

	if err := um.db.Where("hp = ? AND password = ?", hp, password).Find(&res).Error; err != nil {
		log.Println("Terjadi error saat create user", err.Error())
		return User{}, err
	}
	if res.HP == "" || res.Password == "" {
		log.Println("Data tidak ditemukan")
		return User{}, errors.New("data tidak ditemukan")
	}

	return res, nil
}

func (um *UserModel) GetAllUser() ([]User, error) {
	res := []User{}

	if err := um.db.Select("hp, name, id").Find(&res).Error; err != nil {
		log.Println("Terjadi error saat select user ", err.Error())
		return nil, err
	}

	return res, nil
}

// func (m *UserModel) UpdateUser(id uint) (User,error) {
// 	var updatedUser User
// 	if err := m.db.First(&updatedUser, id).Error; err != nil {
// 		log.Println("Terjadi error saat select user ", err.Error())
// 		return User{}, err
// 	}

// 	updatedUser.Name = User.Name
// 	updatedUser.HP = User.HP
// 	updatedUser.Password = User.Password

// 	if err := m.db.Save(&updatedUser).Error; err != nil {
// 		return User{}, err
// 	}

// 	return User{}, nil
// }
// func (um *UserModel) UpdateUser() ([]User, error) {
// 	res := []User{}

// 	if err := um.db.Select("hp, name, id").Find(&res).Error; err != nil {
// 		log.Println("Terjadi error saat select user ", err.Error())
// 		return nil, err
// 	}

// 	return res, nil
// }
