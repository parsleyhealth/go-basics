package main

import (
	"fmt"
)

func main() {
	// Go has multiple ways of declaring and initializing variables

	// PRIMITIVE TYPES
	// Int: int8  int16  int32/rune  int64
	// Uint: uint8/byte uint16 uint32 uint64 uintptr
	// Float: float32 float64
	// Boolean: bool
	// String: string
	// Complex: complex64 complex128

	// Classic
	var first string = "first string"
	var second int = 2
	var third bool = false

	fmt.Println("Classic variables")
	fmt.Println(first)
	fmt.Println(second)
	fmt.Println(third)

	// Shorthand
	fourth := "fourth string"
	fifth := 5
	sixth := true

	fmt.Println("Shorthand variables")
	fmt.Println(fourth)
	fmt.Println(fifth)
	fmt.Println(sixth)

	// Default empty values (Zero types)
	var noString string
	var noInt int
	var noBool bool

	fmt.Println("\nDefault empty values")
	fmt.Println(noString)
	fmt.Println(noInt)
	fmt.Println(noBool)
}
