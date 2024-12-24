package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Marlliton/go/crud-com-auth-jwt/internal/dto"
	"github.com/Marlliton/go/crud-com-auth-jwt/internal/entity"
	"github.com/Marlliton/go/crud-com-auth-jwt/internal/infra/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"
)

type UserHnadler struct {
	UserDB database.UserInterface
}

func NewUserHandler(userDB database.UserInterface) *UserHnadler {
	return &UserHnadler{
		UserDB: userDB,
	}
}

func (h *UserHnadler) Login(w http.ResponseWriter, r *http.Request) {
	jwt := r.Context().Value("jwt").(*jwtauth.JWTAuth)
	jwtExpiresIn := r.Context().Value("jwtExpiresIn").(int)

	var userInputJWT dto.UserLoginInput
	if err := json.NewDecoder(r.Body).Decode(&userInputJWT); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	u, err := h.UserDB.FindByEmail(userInputJWT.Email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if !u.ValidatePassword(userInputJWT.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	_, tokenString, _ := jwt.Encode(map[string]interface{}{
		"sub": u.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(jwtExpiresIn)).Unix(),
	})

	accessToken :=
		struct {
			AccessToken string `json:"access_token"`
		}{
			AccessToken: tokenString,
		}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}

func (h *UserHnadler) Create(w http.ResponseWriter, r *http.Request) {

	var userInput dto.CreateUserInput
	if err := json.NewDecoder(r.Body).Decode(&userInput); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, errs := entity.NewUser(userInput.Name, userInput.Email, userInput.Password)
	if errs != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"sucess": false,
			"errors": errs,
		})
		return
	}

	if err := h.UserDB.Create(u); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *UserHnadler) FindByEmail(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "email")
	if email == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := h.UserDB.FindByEmail(email)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(u)
}
