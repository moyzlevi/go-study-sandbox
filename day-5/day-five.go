package main

import (
	"fmt"
	"runtime"
)

func main() {
	n := 3
	lim := -1

	if v := 3 - n; v < lim {
		fmt.Println(v)
	} else {
		fmt.Println("else: ", v)
	}

	fmt.Println("This OS is: ")
	
	os := runtime.GOOS
	switch os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default :
		fmt.Printf("%s. \n", os)
	}
}
