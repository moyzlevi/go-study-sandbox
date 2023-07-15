package main

import (
	"fmt"
)

type Vertex struct {
	X, Y int
}

func main() {
	// struct literals
	v1 := Vertex{1, 2}
	v2 := Vertex{Y: 1}
	v3 := Vertex{}
	p := &Vertex{}

	fmt.Println(v1, v2, v3, p)

	//arrays
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// slices: dinamic view of arrays
	var s []int = primes[1:4]
	fmt.Println(s)

	// slices are like references to arrays
	names := [4]string{
		"John",
		"Moy",
		"Julia",
		"Marcos",
	}

	fmt.Println(names)

	slice1 := names[0:2]
	slice2 := names[1:3]
	fmt.Println(slice1, slice2)

	slice2[0] = "XXX"
	fmt.Println(slice1, slice2)
	fmt.Println(names)

	// slice literals
	sliceLiteral := []bool{true, false, true}
	fmt.Println(sliceLiteral)

	sliceLiteralWithStrucLiteral := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
	}
	fmt.Println(sliceLiteralWithStrucLiteral)

	// slice defaults
	sliceDefault := []int{2, 3, 4, 5, 6, 7, 8}
	fmt.Println(sliceDefault)
	
	sliceDefault = sliceDefault[:]
	fmt.Println(sliceDefault)
	
	sliceDefault = sliceDefault[1:4]
	fmt.Println(sliceDefault)
	
	sliceDefault = sliceDefault[:2]
	fmt.Println(sliceDefault)

	sliceDefault = sliceDefault[1:]
	fmt.Println(sliceDefault)
	
	sliceDefault = sliceDefault[:]
	fmt.Println(sliceDefault)

	
}
