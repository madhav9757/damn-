package main

import "fmt"

func runArraysDemo() {
	fmt.Println("\n--- ARRAY DEMO ---")

	// ARRAY (fixed size)
	var numbers [5]int = [5]int{1, 2, 3, 4, 5}

	fmt.Println("Array:", numbers)
	fmt.Println("First element:", numbers[0])
	fmt.Println("Length of array:", len(numbers))
}
