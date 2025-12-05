package main

import "fmt"

// greet(): Simple function with no parameters
func greet() {
	fmt.Println("Welcome to Go programming!")
}

// add(): Function with parameters, no return value
func add(a int, b int) {
	fmt.Println("Sum:", a+b)
}

// multiply(): Function with parameters and a return value
func multiply(a int, b int) int {
	return a * b
}

func main() {
	// Calling simple function
	greet()

	// Calling add() function
	add(5, 3)

	// Calling multiply() function and storing the result
	result := multiply(4, 6)
	fmt.Println("Multiplication:", result)
}

/*
------------------- EXPLANATION -------------------

1. func greet() { ... }
   - A function with no input, no output.
   - Just prints a message.

2. func add(a int, b int) { ... }
   - Function that takes 2 integers as input
   - Does not return anything, only prints sum

3. func multiply(a int, b int) int { ... }
   - Function with input and output
   - Returns an integer value using 'return' keyword
   - We can store it in a variable

4. Calling functions in main()
   - Simply write function name with () and parameters if needed
   - Example: greet(), add(5,3), multiply(4,6)

---------------------------------------------------
*/
