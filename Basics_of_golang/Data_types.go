package main

import "fmt"

func main() {
	// Boolean
	var b bool = true
	fmt.Println("Boolean value:", b)

	// Numeric
	var i int = 42
	var f float64 = 3.14
	var c complex128 = complex(1, 2)
	fmt.Println("Integer value:", i)
	fmt.Println("Float value:", f)
	fmt.Println("Complex value:", c)

	// String
	var s string = "Hello, world!"
	fmt.Println("String value:", s)
}
