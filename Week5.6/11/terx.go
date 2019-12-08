package main

import "fmt"

func main() {
	B1 := [4]string{"A", "B", "C", "D"}
	fmt.Println(B1)

	x := B1[0:2]
	fmt.Println(x)

	y := B1[0:2]
	fmt.Println(y)

	z := B1[0:2]
	fmt.Println(z)

}
