package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	fmt.Println("abs was called!")
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type MyString string

func (s MyString) capsLock() MyString {
	return s+" haha";
}

func main() {
	// methods 	
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	var maString MyString = "hello";
	fmt.Println(maString.capsLock());
}