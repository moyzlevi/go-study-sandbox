package main

import "fmt"

func main() {
	// lenght: items in the slice
	// capacity: items in the underlying slice

	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	s = s[:3]
	printSlice(s)

	s = s[:4]
	printSlice(s)

	// jumping the firsts offsets change the capacity of the slice
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}