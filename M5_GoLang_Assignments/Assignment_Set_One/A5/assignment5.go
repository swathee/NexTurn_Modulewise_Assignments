//Exercise 5: Climate Data Analysis

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// City struct to hold climate data
type City struct {
	Name        string
	Temperature float64
	Rainfall    float64
}

func main() {

	cities := []City{
		{"Singapore", 22.5, 120.0},
		{"America", 24.0, 50.0},
		{"Canada", 15.0, 150.0},
		{"Russia", 20.0, 80.0},
		{"Mumbai", 30.0, 250.0},
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\nClimate Data Analysis System")
		fmt.Println("1. Display All Cities")
		fmt.Println("2. City with Highest and Lowest Temperature")
		fmt.Println("3. Average Rainfall")
		fmt.Println("4. Filter Cities by Rainfall Threshold")
		fmt.Println("5. Search by City Name")
		fmt.Println("6. Exit")
		fmt.Print("Choose an option: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			displayCities(cities)
		case "2":
			highest, lowest := findTemperatureExtremes(cities)
			fmt.Printf("City with the Highest Temperature: %s (%.2f°C)\n", highest.Name, highest.Temperature)
			fmt.Printf("City with the Lowest Temperature: %s (%.2f°C)\n", lowest.Name, lowest.Temperature)
		case "3":
			avgRainfall := calculateAverageRainfall(cities)
			fmt.Printf("Average Rainfall Across Cities: %.2f mm\n", avgRainfall)
		case "4":
			fmt.Print("Enter Rainfall Threshold: ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			threshold, err := strconv.ParseFloat(input, 64)
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid number.")
				continue
			}
			filterCitiesByRainfall(cities, threshold)
		case "5":
			fmt.Print("Enter City Name: ")
			cityName, _ := reader.ReadString('\n')
			cityName = strings.TrimSpace(cityName)
			searchCity(cities, cityName)
		case "6":
			fmt.Println("Exiting the system. Thank you!")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}

// Function to display all cities and their data
func displayCities(cities []City) {
	fmt.Println("\nCity Data:")
	for _, city := range cities {
		fmt.Printf("- %s: Temperature = %.2f°C, Rainfall = %.2f mm\n", city.Name, city.Temperature, city.Rainfall)
	}
}

// Function to find the city with the highest and lowest temperatures
func findTemperatureExtremes(cities []City) (highest, lowest City) {
	highest, lowest = cities[0], cities[0]
	for _, city := range cities {
		if city.Temperature > highest.Temperature {
			highest = city
		}
		if city.Temperature < lowest.Temperature {
			lowest = city
		}
	}
	return
}

// Function to calculate the average rainfall across all cities
func calculateAverageRainfall(cities []City) float64 {
	totalRainfall := 0.0
	for _, city := range cities {
		totalRainfall += city.Rainfall
	}
	return totalRainfall / float64(len(cities))
}

// Function to filter and display cities with rainfall above a certain threshold
func filterCitiesByRainfall(cities []City, threshold float64) {
	fmt.Printf("\nCities with Rainfall Above %.2f mm:\n", threshold)
	found := false
	for _, city := range cities {
		if city.Rainfall > threshold {
			fmt.Printf("- %s: %.2f mm\n", city.Name, city.Rainfall)
			found = true
		}
	}
	if !found {
		fmt.Println("No cities found with rainfall above the threshold.")
	}
}

// Function to search for a city by name
func searchCity(cities []City, cityName string) {
	for _, city := range cities {
		if strings.EqualFold(city.Name, cityName) {
			fmt.Printf("\nCity Found: %s\nTemperature: %.2f°C\nRainfall: %.2f mm\n", city.Name, city.Temperature, city.Rainfall)
			return
		}
	}
	fmt.Println("City not found. Please check the name and try again.")
}
