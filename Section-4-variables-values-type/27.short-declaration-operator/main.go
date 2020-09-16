package main

import "fmt"

func main() {
	x := 42 // ":=" is the gopher and represents de short declaration
	fmt.Println(x)
	x = 99
	fmt.Println(x)
	y := 100 + 24
	fmt.Println(y)

	anotherExample()
}

func anotherExample() {
	x := 42 + 7
	y := "james Bond"
	fmt.Println(x)
	fmt.Println(y)
	x = 50
	fmt.Println(x)
}
