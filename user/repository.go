package user

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return Repository{db: db}
}

func (repository *Repository) Register(user User) (err error) {
	query := repository.db.Table("assignment2").Create(&user)

	if query.Error != nil {
		err = fmt.Errorf("error : %s", query.Error)
		return
	}
	query.Order("ID").Last(&user)

	return
}

func (repository *Repository) Login(user User) (getUser User, err error) {
	log.Println("user2", user)
	query := repository.db.Table("assignment2").Where("username", user.Username).Where("password", user.Password).First(&getUser)

	if query.Error != nil {
		err = fmt.Errorf("error : %s", query.Error)
		return
	}

	log.Println("hit sini2")

	return
}
