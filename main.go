package main

import "fmt"

func main() {
	a := [4]int{1, 2, 3, 4}
	reverse(&a)
	fmt.Println(a)
}

func reverse(arr *[4]int) {
	// without dereference we get current value of array on every iteration
	// because arr is pointer(reference), but *arr is a value that range uses once
	for index, value := range *arr {
		arr[len(arr) - 1 - index] = value
	}
}

func double(num *int) {
	*num = *num * 2
}
