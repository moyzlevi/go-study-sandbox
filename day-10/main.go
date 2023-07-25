package main

import (
	"fmt"
	"math")

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
var m map[string]Vertex

var mapLiteral = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

/*
var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}
*/

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

type Vertex struct {
	Lat, Long float64
}

func main() {
	var s []int
	printSlice(s)

	// Append on nil slices
	s = append(s, 0)
	printSlice(s)

	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)

	// Range
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	// skip index assignment
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}

	// maps
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	fmt.Println(mapLiteral)

	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

	// func values
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}