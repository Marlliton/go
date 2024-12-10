package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func newProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/go")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	product := newProduct("notebook3", 1899.00)
	err = insertProduct(db, product)
	if err != nil {
		panic(err)
	}

	product.Price = 100
	err = updateProduct(db, product)
	if err != nil {
		panic(err)
	}

	// p, err := selectProduct(db, product.ID)
	// if err != nil {
	// 	panic(err)
	// }
	//
	// fmt.Printf("Product %v, possui o preço de %.2f", p.Name, p.Price)

	products, err := selectProducts(db)
	if err != nil {
		panic(err)
	}

	for _, product := range products {
		fmt.Printf("Product: %v, possui o preço de %.2f\n", product.Name, product.Price)
	}

	err = deleteProduct(db, product.ID)
	if err != nil {
		panic(err)
	}
}

func deleteProduct(db *sql.DB, id string) error {
	stmt, err := db.Prepare("delete from products where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

func selectProducts(db *sql.DB) ([]Product, error) {
	rows, err := db.Query("select id, name, price from products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var products []Product

	for rows.Next() {
		var prod Product
		err = rows.Scan(&prod.ID, &prod.Name, &prod.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, prod)
	}

	return products, nil
}

func selectProduct(db *sql.DB, id string) (*Product, error) {
	// NOTE: Preparação para proteger a inseção do dados contra sql injection
	stmt, err := db.Prepare("select id, name, price from products where id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var prod Product

	err = stmt.QueryRow(id).Scan(&prod.ID, &prod.Name, &prod.Price)
	if err != nil {
		return nil, err
	}

	return &prod, nil
}

func updateProduct(db *sql.DB, prod *Product) error {
	// NOTE: Preparação para proteger a inseção do dados contra sql injection
	stmt, err := db.Prepare("update products set name = ?, price = ? where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(prod.Name, prod.Price, prod.ID)
	if err != nil {
		return err
	}

	return nil
}

func insertProduct(db *sql.DB, product *Product) error {
	// NOTE: Preparação para proteger a inseção do dados contra sql injection
	stmt, err := db.Prepare("insert into products(id, name, price) values(?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(product.ID, product.Name, product.Price)
	if err != nil {
		return err
	}

	return nil
}
