package main

import "fmt"

var x string
var y int
var z float64
var w bool

func main() {
	fmt.Printf("Zero velue of \"%T\" -> %v\n", x, x)
	fmt.Printf("Zero velue of \"%T\" -> %v\n", y, y)
	fmt.Printf("Zero velue of \"%T\" -> %v\n", z, z)
	fmt.Printf("Zero velue of \"%T\" -> %v\n", w, w)
}
