package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Marlliton/go/sqlc/internal/db"
)

type CourseDB struct {
	dbConn *sql.DB
	*db.Queries
}

func NewCourseDB(dbConn *sql.DB) *CourseDB {
	return &CourseDB{dbConn: dbConn, Queries: db.New(dbConn)}
}

func (c *CourseDB) callTrasacao(ctx context.Context, fn func(*db.Queries) error) error {
	transacao, err := c.dbConn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	query := db.New(transacao) // NOTE: Transação prota para uso
	// NOTE: Call Transação
	err = fn(query)

	if err != nil {
		// NOTE: Tentativa de fazer o rollback case aconteça algum erro ao chamar fn
		// se der erro ao fazer o rollback retornamos o mesmo
		if errRB := transacao.Rollback(); errRB != nil {
			return fmt.Errorf("error on rollback: %v, original error: %w", errRB, err)
		}

		return err
	}

	return transacao.Commit()
}

type (
	CourseParams struct {
		ID          string
		Name        string
		Description sql.NullString
	}

	CategoryParams struct {
		ID          string
		Name        string
		Description sql.NullString
	}
)

func (c *CourseDB) CreateCouseAndCategory(
	ctx context.Context, argsCategory CategoryParams, argsCourse CourseParams,
) error {

	err := c.callTrasacao(ctx, func(q *db.Queries) error {

		var err error
		err = q.CreateCategory(ctx, db.CreateCategoryParams{
			ID:          argsCategory.ID,
			Name:        argsCategory.Name,
			Description: argsCategory.Description,
		})
		if err != nil {
			return err
		}

		err = q.CreateCourse(ctx, db.CreateCourseParams{
			ID:          argsCourse.ID,
			Name:        argsCourse.Name,
			Description: argsCourse.Description,
		})
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func main() {
	ctx := context.Background()
	conn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/sqlc")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	queries := db.New(conn)

	// NOTE: Realizando transação para salvar tudo ou nada.
}
