package main

import "fmt"

func main() {
	B1 := [6]int{0, 1, 2, 3, 4, 5}
	fmt.Println(B1[1])
	B1[1] = 10
	fmt.Println(B1[1])
	length := len(B1)
	fmt.Println("length =", length)

}
