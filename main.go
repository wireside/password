package main

import "fmt"

func main() {
	a := 5
	var pointerA *int = &a // "*" before type name means pointer to the variable with that type
	// pointerA := &a
	
	fmt.Println(pointerA) // prints the address in memory for "a" variable
}

func double(pointer *int) {

}
