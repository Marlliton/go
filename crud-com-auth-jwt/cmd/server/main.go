package main

import (
	"net/http"

	"github.com/Marlliton/go/crud-com-auth-jwt/configs"
	"github.com/Marlliton/go/crud-com-auth-jwt/internal/entity"
	"github.com/Marlliton/go/crud-com-auth-jwt/internal/infra/database"
	"github.com/Marlliton/go/crud-com-auth-jwt/internal/infra/webserver/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	config := configs.LoadConfig(".")

	db, err := gorm.Open(sqlite.Open("dev.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})
	productDB := database.NewProductDB(db)
	productHnadler := handlers.NewProductHandler(productDB)

	userDB := database.NewUserDB(db)
	userHandler := handlers.NewUserHandler(userDB, *config.TokenJWTAuth, config.JWTExpiresIn)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/products", productHnadler.CreateProduct)
	r.Get("/products/{id}", productHnadler.GetProduct)
	r.Get("/products", productHnadler.GetProducts)
	r.Put("/products/{id}", productHnadler.UpdateProduct)
	r.Delete("/products/{id}", productHnadler.DeleteProduct)

	r.Post("/users", userHandler.Create)
	r.Post("/users/login", userHandler.Login)

	http.ListenAndServe(":8000", r)
}
