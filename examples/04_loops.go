package main

import (
	"fmt"
)

func main() {
	// The only loop in Go is the For loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// Mimic the while loop
	c := 0
	for c < 5 {
		fmt.Println(c)
		c++
	}

	// Iterating over a list
	groceries := []string{
		"Potatoes",
		"Tomatoes",
		"Onions",
		"Funyuns",
	}

	for index, groceryItem := range groceries {
		// Formatting verbs @ https://pkg.go.dev/fmt
		fmt.Printf("%d. %s", index, groceryItem)
		fmt.Println()
	}
}
