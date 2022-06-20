package main

import (
	"database/sql"
	"log"

	db2 "github.com/gtkpad/arquitetura-hexagonal-go/adapters/db"
	"github.com/gtkpad/arquitetura-hexagonal-go/application"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", "sqlite.db")
	productDbAdapter := db2.NewProductDb(db)
	productService := application.NewProductService(productDbAdapter)
	product, err := productService.Create("Product Exemple", 30)

	if err != nil {
		log.Fatal((err.Error()))
	}

	productService.Enable(product)
}