package main

import "fmt"

var y = 42
var a int
var b string = "James Bond"

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Printf("%b\n", y)
	fmt.Printf("%x\n", y)
	fmt.Printf("%#x\n", y)
	y = 911
	fmt.Printf("%#x\n", y)
	fmt.Printf("%#x\t%b\t%x\n\n", y, y, y)

	fmt.Println("-------------------------------------")
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%#v\n", a)
	fmt.Printf("%#v\n", b)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\t%T\n", a, b)

	fmt.Println("-------------------------------------")
	s := fmt.Sprint(a, " something more ", b)
	fmt.Println(s)
	s2 := fmt.Sprintf("%v\t%T\t%T\n", "to pass in", a, b)
	fmt.Println(s2)

}

// Difference between these funcs in the "fmt" package

// Group #1 - general printing to standard out
// Print
// Printf
// Println

// Group #2 - printing to a string which you can assign to a variable
// Sprint
// Sprintf
// Sprintln

// Group #3 - printing to a file or a web server's response
// Fprint
// Fprintf
// Fprintln
