package main

import (
	"fmt"
)

// Definition of our struct
type Person struct {
	Name   string
	Height float32
}

// A method of our struct
func (p *Person) Scream() {
	fmt.Printf("And his name is %s!", p.Name)
	fmt.Println()
}

// A function that takes an instance of a struct as an argument
func changeName(p *Person, newName string) {
	p.Name = newName
}

func main() {

	john := Person{
		Name:   "John Doe",
		Height: 1.72,
	}

	fmt.Println(john)

	john.Scream()

	changeName(&john, "John Cena")

	john.Scream()
}
