package main

import "fmt"

type answer int

var x answer

func main() {
	fmt.Printf("%v, %T\n", x, x)

	x = 27
	fmt.Printf("%v, %T\n", x, x)
}
