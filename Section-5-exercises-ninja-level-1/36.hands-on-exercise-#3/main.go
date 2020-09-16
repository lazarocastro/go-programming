package main

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	// fmt.Printf("%v, %T\n", x, x)
	// fmt.Printf("%v, %T\n", y, y)
	// fmt.Printf("%v, %T\n", z, z)

	// My solution
	// s := fmt.Sprint(x, ", ", y, ", ", z)
	// fmt.Println(s)

	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	fmt.Println(s)
}
