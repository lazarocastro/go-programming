package main

import "fmt"

// DECLARE the variable "y"
// ASSIGN the value 43
// declare & assign = initialization
var y = 43 // The scope of "y" would be package level

// DECLARE the is a VARIABEL with the IDENTIFIER "z"
// and that the VARIABEL with the IDENTIFIER "z" is of TYPE int
// ASSIGNS the ZERO VALUE of TYPE int to "z"
// false for boolean, 0 for integers, 0.0 for floats, "" for strings,
// and nil for pointers, functions, interfaces, slices, channels, and maps.
var z int

func main() {
	// short declaration operator
	// DECLARE a variable and ASSIGN a VALUE
	x := 42
	fmt.Println(x)

	fmt.Println(y)

	fmt.Println(z)

	foo()
}

func foo() {
	fmt.Println("Again:", y)
}
