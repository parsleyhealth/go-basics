package main

import (
	"fmt"
)

func main() {
	// Data Structures

	// Arrays
	// Arrays have a fixed size
	volleyBallTeam := [2]string{
		"Joane Joplin",
		"Betty Baxter-Price",
	}

	volleyBallTeam[0] = "Laura Logan"

	fmt.Println(volleyBallTeam)

	// Infering length
	footBallTeam := [...]string{
		"Salvatierra",
		"Solís",
		"Saborío",
	}

	fmt.Println(footBallTeam)

	// Slices
	// Slices are dynamic
	groceries := []string{
		"Potatoes",
		"Tomatoes",
		"Onions",
		"Funyuns",
	}

	fmt.Println(groceries)
}
