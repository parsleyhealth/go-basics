package main

import (
	"fmt"
)

func makeRequest() bool {
	return false
}

func main() {

	// IF STATEMENTS
	height := 130

	// This is the regular way of writing an if statement
	if height < 150 {
		fmt.Println("You can't ride this rollercoaster")
	}

	// We can also do some shorthand for error handling
	if err := makeRequest(); err != false {
		fmt.Println("Error! Request didn't work")
	}

	// SWITCH STATEMENTS
	animal := "cow"

	// A standard Switch
	switch animal {
	case "dog":
		fmt.Println("Bark bark!")
	case "cat":
		fmt.Println("Miau")
	case "cow":
		fmt.Println("Moooo")
	default:
		fmt.Println("ALIENS!")
	}

	// Switch with no starting value
	i := 8
	switch {
	case i > 10:
		fmt.Println("Greater than 10")
	case i < 10:
		fmt.Println("Less than 10")
	default:
		fmt.Println("Is 10")
	}

	i = 12

	// Fallthrough keyword
	switch {
	case i != 10:
		fmt.Println("Does not equal 10")
		fallthrough
	case i < 10:
		fmt.Println("Less than 10")
	case i > 10:
		fmt.Println("Greater than 10")
	default:
		fmt.Println("Is 10")
	}
}
