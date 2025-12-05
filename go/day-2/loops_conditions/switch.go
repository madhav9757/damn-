package main

import "fmt"

func runSwitchDemo() {
	fmt.Println("\n--- SWITCH CASE DEMO ---")

	day := 3

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Another day")
	}
}
