package main

import "fmt"

func main() {
	// ------------------------------
	// 1. Normal variable declaration
	// ------------------------------
	var name string = "Madhav"

	// ----------------------------------------------
	// 2. Type inference (Go auto-detects the datatype)
	// ----------------------------------------------
	var age = 18 // automatically becomes int

	// ----------------------------------------------
	// 3. Short variable declaration (most commonly used)
	// ----------------------------------------------
	city := "Dehradun"

	// -------------------
	// 4. Boolean variable
	// -------------------
	alive := true

	// ----------------
	// 5. Float variable
	// ----------------
	height := 5.9

	// --------------------
	// Printing all values
	// --------------------
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("City:", city)
	fmt.Println("Alive:", alive)
	fmt.Println("Height:", height)

	// ---------------------------
	// Taking input from the user
	// ---------------------------
	var inputName string
	fmt.Print("Enter your name: ")
	fmt.Scan(&inputName) // & gives memory address so Scan can store input
	fmt.Println("Hello,", inputName)
}
