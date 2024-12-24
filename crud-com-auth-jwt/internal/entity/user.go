package entity

import (
	"github.com/Marlliton/validator"
	"github.com/Marlliton/validator/rules"
	"github.com/Marlliton/validator/validator_error"
	"golang.org/x/crypto/bcrypt"

	"github.com/Marlliton/go/crud-com-auth-jwt/pkg/entity"
)

type User struct {
	ID       entity.ID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"-"` // Não devolve a senha para o usuário na request
}

func NewUser(name, email, password string) (*User, []*validator_error.ValidatorError) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, []*validator_error.ValidatorError{{Field: "Password", Message: "Failed to hash password"}}
	}

	user := User{
		ID:       entity.NewID(),
		Name:     name,
		Email:    email,
		Password: string(hash),
	}

	errs, ok := user.Validate()
	if !ok {
		return nil, errs
	}

	return &user, nil
}

func (u *User) Validate() ([]*validator_error.ValidatorError, bool) {
	v := validator.New()

	v.Add("Name", rules.Rules{
		rules.Required(),
		rules.MinLength(3),
		rules.MaxLength(80),
	})
	v.Add("Email", rules.Rules{
		rules.Required(),
		rules.ValidEmail(),
	})
	v.Add("Password", rules.Rules{
		rules.Required(),
	})
	errs := v.Validate(*u)
	if len(errs) == 0 {
		return nil, true
	}

	var validationErr []*validator_error.ValidatorError
	for _, err := range errs {
		validationErr = append(validationErr, err)
	}

	return validationErr, false
}

func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
