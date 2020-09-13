package main

import "fmt"

var y = 42

// DECLARED the VARIABLE with the IDENTIFIER "z"
// is of TYPE string
// and I ASSIGN the VALUE "Shaken, not stirred"

var z string = "Shaken, not stirred"
var x string = `James Bond say,
"Shaken, 
not stirred"`

// This is a STATIC programming language
// a VARIABLE is DECLARED to hold a VALUE of a certain TYPE
// not a DYNAMIC programming language

func main() {
	fmt.Printf("%v - %T\n\n", y, y)
	fmt.Printf("%v - %T\n\n", z, z)
	fmt.Printf("%v - %T\n\n", x, x)
}
