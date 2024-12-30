//Exercise 3: Inventory Management System

package main

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

// Product struct representing a product in the inventory
type Product struct {
	ID    int
	Name  string
	Price float64
	Stock int
}

// Inventory slice to hold all products
var inventory []Product

// AddProduct adds a new product to the inventory
func AddProduct(id int, name string, price float64, stock int) error {

	for _, product := range inventory {
		if product.ID == id {
			return errors.New("product ID already exists")
		}
	}

	inventory = append(inventory, Product{ID: id, Name: name, Price: price, Stock: stock})
	return nil
}

// UpdateStock updates the stock of a specific product by ID
func UpdateStock(id int, newStock int) error {
	if newStock < 0 {
		return errors.New("stock cannot be negative")
	}
	for i, product := range inventory {
		if product.ID == id {
			inventory[i].Stock = newStock
			return nil
		}
	}
	return errors.New("product not found")
}

// SearchProduct searches for a product by ID or name
func SearchProduct(query string) (*Product, error) {
	for _, product := range inventory {
		if strings.EqualFold(product.Name, query) || fmt.Sprintf("%d", product.ID) == query {
			return &product, nil
		}
	}
	return nil, errors.New("product not found")
}

// DisplayInventory displays all products in the inventory
func DisplayInventory() {
	fmt.Printf("%-5s %-20s %-10s %-10s\n", "ID", "Name", "Price", "Stock")
	fmt.Println(strings.Repeat("-", 50))
	for _, product := range inventory {
		fmt.Printf("%-5d %-20s %-10.2f %-10d\n", product.ID, product.Name, product.Price, product.Stock)
	}
}

// SortInventory sorts products by a given field (price or stock)
func SortInventory(by string) {
	switch by {
	case "price":
		sort.Slice(inventory, func(i, j int) bool {
			return inventory[i].Price < inventory[j].Price
		})
	case "stock":
		sort.Slice(inventory, func(i, j int) bool {
			return inventory[i].Stock < inventory[j].Stock
		})
	}
}

func main() {
	fmt.Println("Welcome to the Inventory Management System")

	// Add some initial products
	_ = AddProduct(1, "Earrings", 200.00, 30)
	_ = AddProduct(2, "Necklace", 500.00, 20)
	_ = AddProduct(3, "Rings", 120.00, 10)
	_ = AddProduct(4, "Bracelets", 150.00, 50)

	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Add Product")
		fmt.Println("2. Update Stock")
		fmt.Println("3. Search Product")
		fmt.Println("4. Display Inventory")
		fmt.Println("5. Sort Inventory by Price")
		fmt.Println("6. Sort Inventory by Stock")
		fmt.Println("7. Exit")
		fmt.Print("Choose an option: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var id, stock int
			var name string
			var price float64
			fmt.Print("Enter Product ID: ")
			fmt.Scan(&id)
			fmt.Print("Enter Product Name: ")
			fmt.Scan(&name)
			fmt.Print("Enter Product Price: ")
			fmt.Scan(&price)
			fmt.Print("Enter Product Stock: ")
			fmt.Scan(&stock)
			if err := AddProduct(id, name, price, stock); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Product added successfully.")
			}
		case 2:
			var id, newStock int
			fmt.Print("Enter Product ID: ")
			fmt.Scan(&id)
			fmt.Print("Enter New Stock: ")
			fmt.Scan(&newStock)
			if err := UpdateStock(id, newStock); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Stock updated successfully.")
			}
		case 3:
			var query string
			fmt.Print("Enter Product Name or ID: ")
			fmt.Scan(&query)
			if product, err := SearchProduct(query); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("Found: %+v\n", *product)
			}
		case 4:
			DisplayInventory()
		case 5:
			SortInventory("price")
			fmt.Println("Inventory sorted by price.")
			DisplayInventory()
		case 6:
			SortInventory("stock")
			fmt.Println("Inventory sorted by stock.")
			DisplayInventory()
		case 7:
			fmt.Println("Exiting... Thank you!")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
