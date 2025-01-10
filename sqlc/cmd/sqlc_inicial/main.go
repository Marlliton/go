package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Marlliton/go/sqlc/internal/db"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

func main() {
	ctx := context.Background()
	conn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/sqlc")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	queries := db.New(conn)

	err = queries.CreateCategory(ctx, db.CreateCategoryParams{
		ID:          uuid.New().String(),
		Name:        "Backend",
		Description: sql.NullString{String: "Backend Description", Valid: true},
	})
	if err != nil {
		panic(err)
	}

	categories, err := queries.ListCategories(ctx)
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		fmt.Println(category.ID, category.Name, category.Description.String)
	}
}
