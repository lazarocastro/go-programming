package main

import "fmt"

// Declare variable with regular "int" type
var a int

// Create a variable with the "hotdog int" type
type hotdog int

var b hotdog

func main() {
	a = 42
	b = 43
	fmt.Printf("%v %T\n", a, a)
	fmt.Printf("%v %T\n", b, b)

	// But if you try to assing "a = b" its not allowed
	// Because a is "int" but type is "hotdog"
	// You need to convert to "int" to assing "b" inside "a"
	a = int(b)
	fmt.Printf("%v %T\n", a, a)
}
