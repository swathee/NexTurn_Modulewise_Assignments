///Exercise 1: Employee Management System

package main

import (
	"errors"
	"fmt"
	"strings"
)

// Employee struct to hold employee data
type Employee struct {
	ID         int
	Name       string
	Age        int
	Department string
}

// Constants for departments
const (
	HR    = "HR"
	IT    = "IT"
	Sales = "Sales"
)

// In-memory employee database
var employees []Employee

// AddEmployee adds a new employee to the database
func AddEmployee(id int, name string, age int, department string) error {
	// Validate ID uniqueness
	for _, emp := range employees {
		if emp.ID == id {
			return errors.New("employee with this ID already exists")
		}
	}

	// Validate Age
	if age <= 18 {
		return errors.New("employee age must be greater than 18")
	}

	// Add employee to the database
	employee := Employee{
		ID:         id,
		Name:       name,
		Age:        age,
		Department: department,
	}
	employees = append(employees, employee)
	return nil
}

// SearchEmployee searches for an employee by ID or name
func SearchEmployee(searchTerm string) (*Employee, error) {
	for _, emp := range employees {
		if fmt.Sprintf("%d", emp.ID) == searchTerm || strings.EqualFold(emp.Name, searchTerm) {
			return &emp, nil
		}
	}
	return nil, errors.New("employee not found")
}

// ListEmployeesByDepartment lists employees in a specific department
func ListEmployeesByDepartment(department string) ([]Employee, error) {
	var filteredEmployees []Employee
	for _, emp := range employees {
		if strings.EqualFold(emp.Department, department) {
			filteredEmployees = append(filteredEmployees, emp)
		}
	}
	if len(filteredEmployees) == 0 {
		return nil, errors.New("no employees found in this department")
	}
	return filteredEmployees, nil
}

// CountEmployees counts the employees in a specific department
func CountEmployees(department string) int {
	count := 0
	for _, emp := range employees {
		if strings.EqualFold(emp.Department, department) {
			count++
		}
	}
	return count
}

func main() {
	// Adding sample employees
	_ = AddEmployee(1, "Remya", 25, IT)
	_ = AddEmployee(2, "Poorna", 30, Sales)
	_ = AddEmployee(3, "Swathee", 22, IT)
	_ = AddEmployee(4, "Mini", 29, HR)

	fmt.Println("=== Employee Management System ===")

	// 1. Add Employee
	fmt.Println("\nAdding Employee...")
	err := AddEmployee(5, "Daya", 17, HR) // Invalid age
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("\nAdding Employee...")
	err1 := AddEmployee(5, "Varshini", 21, Sales) // valid age
	fmt.Println("Successfully Added!")
	if err1 != nil {
		fmt.Println("Error:", err1)
	}

	// 2. Search Employee
	fmt.Println("\nSearching for Employee by ID...")
	emp, err := SearchEmployee("2")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Employee Found: %+v\n", *emp)
	}

	// 3. List Employees by Department
	fmt.Println("\nListing Employees in IT Department...")
	itEmployees, err := ListEmployeesByDepartment(IT)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		for _, e := range itEmployees {
			fmt.Printf("Employee: %+v\n", e)
		}
	}

	// 4. Count Employees
	fmt.Println("\nCounting Employees in HR Department...")
	hrCount := CountEmployees(HR)
	fmt.Printf("Number of Employees in HR: %d\n", hrCount)

	// Bonus: Handling invalid operations
	fmt.Println("\nSearching for a non-existent Employee...")
	_, err = SearchEmployee("100")
	if err != nil {
		fmt.Println("Error:", err)
	}
}
