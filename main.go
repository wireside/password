package main

import "fmt"

func main() {
	a := 5
	var pointerA *int = &a // "*" before type abbreviation means pointer to the variable with that type
	// pointerA := &a
	fmt.Println(pointerA)  // prints the address in memory for "a" variable
	fmt.Println(*pointerA) // * means dereference for pointer,
	// in other words it's way to get the variable and its value from pointer
	
	fmt.Println("Value before:", a)
	
	double(&a)
	
	fmt.Println("Value after:", a)
	
	// with a pointer you can pass the variable as reference, not as value
	// reference types like slice, map, function, chanel passes as reference by default
}

func double(num *int) {
	*num = *num * 2
}
