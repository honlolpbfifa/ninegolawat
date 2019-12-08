package main

import "fmt"

func main() {
	x := make([]int, 5)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	y := make([]int, 5, 10)
	fmt.Println(y)
	fmt.Println(len(y))
	fmt.Println(cap(y))

	z := make([]int, 2, 3)
	fmt.Println(z)
	fmt.Println(cap(z))
	fmt.Println(len(z))

}
