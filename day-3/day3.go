package main

import (
	"fmt"
	"math/cmplx"
)

// rune: alias for int32
// float32, float64
// complex64, complex128
// int, int8, int16, int32, int64
// uint, uint8, uint16, uint32, uint64, uintptr

var (
	ToBe bool = false
	MaxInt uint64 = 1<< 64-1
	z complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	// normal var
	var i int = 2

	//ready assign :=
	autoAssign := 3
	fmt.Println("Auto assign: ", autoAssign, i)

	fmt.Printf("type: %T value %v \n", ToBe, ToBe)
	fmt.Printf("type: %T value %v \n", MaxInt, MaxInt)
	fmt.Printf("type: %T value %v \n", z, z)
}

func basicType() {
	var i bool = false
	var x string = "hello"
	var l =3

	fmt.Println(i, x, l)
}
