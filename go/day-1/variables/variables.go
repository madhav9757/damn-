package main

import "fmt"

func main() {

	// 1. Normal variable declaration
	var name string = "Madhav"

	// 2. Type inference (Go automatically detects type)
	var age = 18

	// 3. Short variable declaration (most commonly used)
	city := "Dehradun"

	// 4. Boolean variable
	alive := true

	// 5. Float variable
	height := 5.9

	// Print all variables
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("City:", city)
	fmt.Println("Alive:", alive)
	fmt.Println("Height:", height)

	// -------------------
	// Taking input from user
	var inputName string
	fmt.Print("Enter your name: ")
	fmt.Scan(&inputName) // &inputName -> memory address of variable
	fmt.Println("Hello,", inputName)
}

/*
------------------- EXPLANATION -------------------

1. var name string = "Madhav"
   - Declares a variable 'name' of type string and assigns a value.

2. var age = 18
   - Declares 'age' without specifying type. Go automatically detects int.

3. city := "Dehradun"
   - Short-hand declaration, most used in Go. Type is inferred.

4. Boolean & Float
   - alive := true  -> boolean
   - height := 5.9 -> float64

5. fmt.Println(...)
   - Prints variables to the screen

6. User Input
   - fmt.Print("...") prints message without newline
   - fmt.Scan(&inputName) waits for user input
   - '&' is used because Go passes the memory address to store input

---------------------------------------------------
*/
