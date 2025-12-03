package main

import "fmt"

// main(): This is the starting point of every Go program
func main() {

	// Print a message on screen
	fmt.Println("Hello, Go! This is my first program.")
}

/*
------------------- EXPLANATION -------------------

1. package main
   - Every Go program starts with a package.
   - 'main' means this is an executable program.

2. import "fmt"
   - 'fmt' is a Go package which gives us functions to print text.
   - Example: fmt.Println()

3. func main() { }
   - main() is the first function Go runs.
   - Whatever you write inside main() gets executed.

4. fmt.Println("...")
   - This prints the text on the screen with a new line.

---------------------------------------------------
*/
