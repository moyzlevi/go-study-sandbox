package main

import (
	"fmt"
)

// structs are collections of fields
type Vertex struct {
	X int
	Y int
}

func main() {

	// defers are executend when the surrounding function returns
	defer fmt.Println("I'm deferred!")

	// defers are stacked in a last-in-first-out
	for i := 0; i < 3; i++ {
		defer fmt.Println(i)
	}

	// switches can also be like if and elses
	phrase := "helo"

	switch {
	case len(phrase) == 3:
		fmt.Println("Length is 3!")
	case len(phrase) == 4:
		fmt.Println("Length is 4!")
	default:
		fmt.Println("Other!")
	}

	// zero value is nil
	var pointer *int
	fmt.Println(pointer)

	// & operator generates a pointer
	i := 42
	pointer = &i

	fmt.Println(pointer)

	// *operator denotes de pointer's underlying value
	fmt.Println("Underlying: ", *pointer)
	*pointer = 32

	fmt.Println("Underlying value changed: ", i)

	// structs
	fmt.Println(Vertex{1, 2})
	// struct fields are accessed using a dot
	v := Vertex{1,2}
	v.X = 4
	fmt.Println("v.X ",v.X)

	// when whe have a pointer to a struct we can access a field as we were in the underlying value
	v2 := Vertex{1 ,2}
	pointer2 := &v2
	pointer2.X = 1e9
	// or the canonical way
	(*pointer2).Y = -1
	fmt.Println("Underlying value pointer2: ",v2)
}
