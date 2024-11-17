package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"

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
	dsn := "root@tcp(127.0.0.1:3307)/go_crud_products" // Without password
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

// Create a new product in the database
func createProduct(name, description string, price float64) {
	_, err := db.Exec("INSERT INTO products (name, description, price) VALUES (?, ?, ?)", name, description, price)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Product created successfully!")
}

// Get a product by its ID
func getProduct(id int) (*Product, error) {
	var product Product
	err := db.QueryRow("SELECT id, name, description, price FROM products WHERE id = ?", id).Scan(&product.ID, &product.Name, &product.Description, &product.Price)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

// Get all products from the database
func getAllProducts() ([]Product, error) {
	rows, err := db.Query("SELECT id, name, description, price FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

// Update a product by its ID
func updateProduct(id int, name, description string, price float64) {
	_, err := db.Exec("UPDATE products SET name = ?, description = ?, price = ? WHERE id = ?", name, description, price, id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Product updated successfully!")
}

// Delete a product by its ID
func deleteProduct(id int) {
	_, err := db.Exec("DELETE FROM products WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Product deleted successfully!")
}

func showMenu() {
	fmt.Println("\nMenu:")
	fmt.Println("1. View All Products")
	fmt.Println("2. Add New Product")
	fmt.Println("3. Edit Product")
	fmt.Println("4. Delete Product")
	fmt.Println("5. Exit")
}

func main() {
	var choice int

	for {
		showMenu()
		fmt.Print("Choose an option: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			// View all products
			products, err := getAllProducts()
			if err != nil {
				log.Fatal(err)
			}
			if len(products) == 0 {
				fmt.Println("No products found.")
			} else {
				fmt.Println("\nProduct List:")
				for _, product := range products {
					fmt.Printf("ID: %d, Name: %s, Description: %s, Price: %.2f\n", product.ID, product.Name, product.Description, product.Price)
				}
			}
		case 2:
			// Add new product
			var name, description string
			var price float64

			fmt.Print("Enter product name: ")
			fmt.Scanln(&name)

			// Use bufio.Scanner for multi-line input
			fmt.Print("Enter product description: ")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan() // Read the full description until Enter is pressed
			description = scanner.Text()

			fmt.Print("Enter product price: ")
			// Ensure the price is properly captured as a float
			fmt.Scanln(&price)

			createProduct(name, description, price)
		case 3:
			// Edit product
			var id int
			var name, description string
			var price float64

			fmt.Print("Enter product ID to edit: ")
			fmt.Scanln(&id)

			product, err := getProduct(id)
			if err != nil {
				fmt.Println("Error: Product not found!")
			} else {
				fmt.Println("Current Product:", product)
				fmt.Print("Enter new product name: ")
				fmt.Scanln(&name)

				fmt.Print("Enter new product description: ")
				scanner := bufio.NewScanner(os.Stdin)
				scanner.Scan()
				description = scanner.Text()

				fmt.Print("Enter new product price: ")
				fmt.Scanln(&price)

				updateProduct(id, name, description, price)
			}
		case 4:
			// Delete product
			var id int
			fmt.Print("Enter product ID to delete: ")
			fmt.Scanln(&id)

			product, err := getProduct(id)
			if err != nil {
				fmt.Println("Error: Product not found!")
			} else {
				fmt.Println("Product to delete:", product)
				deleteProduct(id)
			}
		case 5:
			// Exit
			fmt.Println("Exiting program...")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
