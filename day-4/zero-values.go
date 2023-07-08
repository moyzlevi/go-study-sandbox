package main

import "fmt"

func main() {
	// default init 
	// var i int
	// var f float64
	// var b bool
	// var s string
	// // fmt.Printf("%v %v %v %q\n", i, f, b, s)

	// type converions
	myVar := 42
	var converted float64 = float64(myVar)
	// fmt.Printf("wtf: %T", converted)

	fmt.Printf("type conversion: %T\n", converted)
}