package main

import "fmt"

func main() {

	// -----------------
	// 1. ARRAY (fixed size)
	var numbers [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array:", numbers)
	fmt.Println("First element:", numbers[0])
	fmt.Println("Length of array:", len(numbers))

	// -----------------
	// 2. SLICE (dynamic size)
	fruits := []string{"Apple", "Banana", "Mango"}
	fmt.Println("Slice:", fruits)

	// Adding an element
	fruits = append(fruits, "Orange")
	fmt.Println("After append:", fruits)

	// Slice of slice
	numbersSlice := numbers[1:4] // elements index 1 to 3
	fmt.Println("Numbers slice:", numbersSlice)

	// -----------------
	// 3. MAP (key-value pairs)
	scores := map[string]int{
		"Madhav": 90,
		"Rahul":  85,
		"Anita":  95,
	}
	fmt.Println("Scores map:", scores)

	// Access value by key
	fmt.Println("Madhav's score:", scores["Madhav"])

	// Add new key-value
	scores["Riya"] = 88
	fmt.Println("After adding Riya:", scores)

	// Delete a key
	delete(scores, "Rahul")
	fmt.Println("After deleting Rahul:", scores)

	// Loop over map
	fmt.Println("Looping over scores:")
	for name, score := range scores {
		fmt.Println(name, ":", score)
	}
}

/*
------------------- EXPLANATION -------------------

1. ARRAY
   - Fixed size collection of same type elements
   - Access by index: numbers[0]
   - len(array) gives length

2. SLICE
   - Dynamic array (most used)
   - Can append elements using append()
   - Can create slice from array using numbers[1:4]

3. MAP
   - Key-value collection
   - Access value by key: scores["Madhav"]
   - Add: scores["NewName"] = value
   - Delete: delete(scores, "Key")
   - Loop with for name, score := range scores

---------------------------------------------------
*/
