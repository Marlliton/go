package database

import (
	"github.com/Marlliton/go/crud-com-auth-jwt/internal/entity"
	"gorm.io/gorm"
)

type User struct {
	DB *gorm.DB
}

func NewUserDB(db *gorm.DB) *User {
	return &User{DB: db}
}

func (u *User) Create(user *entity.User) error {
	return u.DB.Create(user).Error
}

func (u *User) FindByEmail(e string) (*entity.User, error) {
	var user entity.User
	if err := u.DB.Where("email = ?", e).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
