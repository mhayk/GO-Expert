package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}

	// // Create a new product repository
	// productRepository := product.NewProductRepository(db)

	// // Create a new product use case
	// productUseCase := product.NewProductUseCase(productRepository)

	useCase := NewUseCase(db)

	// Get a product by its ID
	product, err := useCase.GetProduct(1)
	if err != nil {
		panic(err)
	}

	// Print the product
	fmt.Println(product.Name)
}
