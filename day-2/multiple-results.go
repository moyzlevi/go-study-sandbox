package main

import "fmt"


var pizza, pasta, apple bool

// naked return
func imNaked(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// two returns
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "swap")
	fmt.Println(a, b)

	fmt.Println(imNaked(17))

	var i int
	fmt.Println(i, pizza, pasta, apple)

	var c, python, java = true, false, "it pays my bills!"
	fmt.Println(c, python, java)
}
