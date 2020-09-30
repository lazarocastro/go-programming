package main

import (
	"errors"
	"fmt"
	"math"
)

// variables in package scope
var x int
var a [5]int

// creating struct
type person struct {
	name string
	age  int
}

func main() {

	// assing variable x
	x = 42
	fmt.Println(x)

	// gopher assing variable or short assingment
	y := 99
	fmt.Println(y)

	// conditionals
	if x > y {
		fmt.Println("x > y")
	} else if x < y {
		fmt.Println("x < y")
	} else {
		fmt.Println("Last else")
	}

	// arrays
	a[2] = 7
	fmt.Println(a)
	b := [5]int{5, 4, 3, 2, 1}
	// You cant use append in arrays
	// b = append(b, 13)
	fmt.Println(b)

	// slices
	c := []int{1, 2, 3, 4, 5}
	// append into slices
	c = append(c, 13)
	fmt.Println(c)

	// maps
	vertices := make(map[string]int)
	vertices["triangle"] = 2
	vertices["square"] = 3
	vertices["dodecagon"] = 12
	fmt.Println("trigles ->", vertices["triangle"])
	fmt.Println(vertices)
	delete(vertices, "square")
	fmt.Println(vertices)

	// loops
	// for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	// while loop
	d := 0
	for d < 5 {
		fmt.Println(d)
		d++
	}
	// iterate over eacho element in an array or slice
	// by using for loop
	arr := []string{"a", "b", "c"}
	for index, value := range arr {
		fmt.Println("index:", index, "value:", value)
	}
	// Do the same thing with a map but you get the key
	// instead os the index
	m := make(map[string]string)
	m["a"] = "alpha"
	m["b"] = "bate"

	for key, value := range m {
		fmt.Println("key:", key, "valeu:", value)
	}

	// call function sum()
	result := sum(2, 5)
	fmt.Println(result)

	//call function sqrt()
	res, err := sqrt(-16)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("sqrt is", res)
	}

	// using struct person
	p := person{name: "Lazaro", age: 39}
	fmt.Println(p)
	// if you want to get a specific field
	// out the struct then you can use dot
	// notation to do that
	fmt.Println(p.age)
	fmt.Println(p.name)

	// points
	point := 7
	fmt.Println(point)
	fmt.Println(&point)

	bar()

}

// creating functions
func sum(x int, y int) int {
	return x + y
}

// function with multiples return values
func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefined for negative numbers")
	}

	return math.Sqrt(x), nil
}

func bar() {
	fmt.Println("This is a test")
}
