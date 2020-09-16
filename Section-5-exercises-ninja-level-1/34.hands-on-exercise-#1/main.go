package main

import "fmt"

func main() {
	a := 42
	b := "James Bond"
	c := true

	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", c, c)
}
