package main

import (
	"fmt"
)

// Anatomy of a function
// 1. First we need to specify the declaration with the func keyword
// 2. Next we specify the name of our function (addition)
// 3. After that we open parentheses to specify or arguments. Arguments
// come with the name first and then the type. An empty parentheses means
// no arguments
// 4. Then, after the parentheses we specify the return value. Leave blank
// if there is no return value
// 5. Finally we open brackets and write our logic.

func addition(num1 int, num2 int) int {
	return num1 + num2
}

// Multiple returns
func addAndMult(num1 int, num2 int) (int, int) {
	add := num1 + num2
	mult := num1 * num2

	return add, mult
}

// Named returns
func whatsInThaBox() (box string) {
	box = "You don't wanna know"

	return
}

// Varidic Functions
func additionPlusPlus(nums ...int) int {
	total := 0

	for _, num := range nums {
		total += num
	}

	return total
}

func main() {
	// Functions
	sum := addition(4, 6)
	fmt.Println(sum)

	sum2, mult := addAndMult(sum, 6)
	fmt.Printf("Sum: %d, Mult: %d", sum2, mult)
	fmt.Println()

	morganFreemanSays := whatsInThaBox()
	fmt.Println(morganFreemanSays)

	sumAll := additionPlusPlus(sum, sum2, mult)
	fmt.Println(sumAll)
}
