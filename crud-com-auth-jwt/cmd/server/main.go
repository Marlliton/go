package main

import (
	"net/http"

	"github.com/Marlliton/go/crud-com-auth-jwt/configs"
	_ "github.com/Marlliton/go/crud-com-auth-jwt/docs"
	"github.com/Marlliton/go/crud-com-auth-jwt/internal/entity"
	"github.com/Marlliton/go/crud-com-auth-jwt/internal/infra/database"
	"github.com/Marlliton/go/crud-com-auth-jwt/internal/infra/webserver/handlers"
	internal_middleware "github.com/Marlliton/go/crud-com-auth-jwt/internal/infra/webserver/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// @title Crud com autenticação
// @version 1.0.0
// @description Crud de produtos com autenticação

// @contact.name Marlliton Souza
// @contact.email marlliton.souza1@gmail.com

// @host localhost:8000
// @basePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

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
	userHandler := handlers.NewUserHandler(userDB)

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.WithValue("jwt", config.TokenJWTAuth))
	r.Use(middleware.WithValue("jwtExpiresIn", config.JWTExpiresIn))
	r.Use(internal_middleware.JSONContentType)

	r.Route("/products", func(r chi.Router) {
		// NOTE: Protegendo rotas de produto
		r.Use(jwtauth.Verifier(config.TokenJWTAuth))
		r.Use(jwtauth.Authenticator)

		r.Post("/", productHnadler.CreateProduct)
		r.Get("/{id}", productHnadler.GetProduct)
		r.Get("/", productHnadler.GetProducts)
		r.Put("/{id}", productHnadler.UpdateProduct)
		r.Delete("/{id}", productHnadler.DeleteProduct)
	})

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))

	r.Route("/users", func(r chi.Router) {
		r.Post("/", userHandler.Create)
		r.Post("/login", userHandler.Login)
	})

	http.ListenAndServe(":8000", r)
}
