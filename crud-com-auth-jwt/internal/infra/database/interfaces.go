package database

import "github.com/Marlliton/go/crud-com-auth-jwt/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
