package main

import "fmt"

func main() {

	// -----------------
	// IF-ELSE CONDITION
	age := 18

	if age < 18 {
		fmt.Println("You are a minor")
	} else if age == 18 {
		fmt.Println("You just became an adult!")
	} else {
		fmt.Println("You are an adult")
	}

	// -----------------
	// SWITCH CASE
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

	// -----------------
	// FOR LOOP (classic)
	fmt.Println("Numbers from 1 to 5:")
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// -----------------
	// FOR LOOP as WHILE (only condition)
	fmt.Println("Counting down from 5:")
	count := 5
	for count > 0 {
		fmt.Println(count)
		count--
	}

	// -----------------
	// FOR LOOP over array
	fruits := []string{"Apple", "Banana", "Mango"}
	fmt.Println("Fruits list:")
	for index, fruit := range fruits {
		fmt.Println(index+1, ":", fruit)
	}
}

/*
------------------- EXPLANATION -------------------

1. IF-ELSE
   - Standard condition check
   - Example: if age < 18 { ... } else if ... else ...

2. SWITCH CASE
   - Checks multiple values of a variable
   - default is executed if no match

3. FOR LOOP
   - Classic for: for i := 1; i <= 5; i++
   - While-style: for count > 0 { count-- }
   - Go doesn't have 'while' keyword, we use 'for' for everything

4. FOR RANGE
   - Loop over arrays/slices/maps
   - range returns index and value

---------------------------------------------------
*/
