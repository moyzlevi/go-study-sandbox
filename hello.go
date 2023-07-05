package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("Hello World!")
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println("The number PI is", math.Pi)
	fmt.Println("Let's add two numbers! 3 + 5 = ", add(3,5) )

	fmt.Println("Let's subtract two numbers! 5 - 2 = ", minus(5,2, "", "") )
}

func add(x int, y int) int {
	return x + y
}


// we can subtract the first types if they're equal the last one
func minus(x, y int, a , b string) int {
	return x - y
}