package main

import "fmt"

func runSlicesDemo() {
	fmt.Println("\n--- SLICE DEMO ---")

	fruits := []string{"Apple", "Banana", "Mango"}
	fmt.Println("Slice:", fruits)

	// Add element
	fruits = append(fruits, "Orange")
	fmt.Println("After append:", fruits)

	// Slice from array example
	numbers := [5]int{1, 2, 3, 4, 5}
	numbersSlice := numbers[1:4]
	fmt.Println("Numbers slice:", numbersSlice)
}
