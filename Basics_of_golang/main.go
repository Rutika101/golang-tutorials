package main

import "fmt"

func main() {
	// Array
	var a [3]int
	a[0] = 1
	a[1] = 2
	a[2] = 3
	fmt.Println("Array values:", a)

	// Slice
	var s []int = []int{4, 5, 6}
	fmt.Println("Slice values:", s)

	// Struct
	type Person struct {
		Name string
		Age  int
	}
	var p Person
	p.Name = "John"
	p.Age = 30
	fmt.Println("Struct value:", p)
}
