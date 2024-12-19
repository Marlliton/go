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
	UserDB       database.UserInterface
	JWT          jwtauth.JWTAuth
	JWTExpiresIn int
}

func NewUserHandler(userDB database.UserInterface, jwt jwtauth.JWTAuth, jwtExpiresIn int) *UserHnadler {
	return &UserHnadler{
		UserDB:       userDB,
		JWT:          jwt,
		JWTExpiresIn: jwtExpiresIn,
	}
}

func (h *UserHnadler) Login(w http.ResponseWriter, r *http.Request) {
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

	_, tokenString, _ := h.JWT.Encode(map[string]interface{}{
		"sub":        u.ID.String(),
		"expires_in": time.Now().Add(time.Second * time.Duration(h.JWTExpiresIn)).Unix(),
	})

	accessToken :=
		struct {
			AccessToken string `json:"access_token"`
		}{
			AccessToken: tokenString,
		}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}

func (h *UserHnadler) Create(w http.ResponseWriter, r *http.Request) {
	var userInput dto.CreateUserInput
	if err := json.NewDecoder(r.Body).Decode(&userInput); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := entity.NewUser(userInput.Name, userInput.Email, userInput.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
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
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(u)
}
