package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// Product struct
type Product struct {
	ID          int
	Name        string
	Description string
	Price       float64
}

func init() {
	// Connect to MySQL (no password)
	var err error
	dsn := "root@tcp(127.0.0.1:3307)/go_crud_products" // Tanpa password
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

func createProduct(name, description string, price float64) {
	_, err := db.Exec("INSERT INTO products (name, description, price) VALUES (?, ?, ?)", name, description, price)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Product created successfully!")
}

func getProduct(id int) (*Product, error) {
	var product Product
	err := db.QueryRow("SELECT id, name, description, price FROM products WHERE id = ?", id).Scan(&product.ID, &product.Name, &product.Description, &product.Price)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func updateProduct(id int, name, description string, price float64) {
	_, err := db.Exec("UPDATE products SET name = ?, description = ?, price = ? WHERE id = ?", name, description, price, id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Product updated successfully!")
}

func deleteProduct(id int) {
	_, err := db.Exec("DELETE FROM products WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Product deleted successfully!")
}

func main() {
	
	var name, description string
	var price float64

	
	fmt.Print("Enter product name: ")
	fmt.Scanln(&name)

	fmt.Print("Enter product description: ")
	fmt.Scanln(&description)

	fmt.Print("Enter product price: ")
	fmt.Scanln(&price)

	
	createProduct(name, description, price)

	
	product, err := getProduct(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Retrieved product:", product)
}
