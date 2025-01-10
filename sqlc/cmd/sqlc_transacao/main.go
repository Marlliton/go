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
