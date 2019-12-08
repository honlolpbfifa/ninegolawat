package main

import "fmt"

func main() {
	B1 := [4]string{"A", "B", "C", "D"}
	fmt.Println(B1)

	x := B1[0:2]
	fmt.Println(x)

	y := B1[2:4]
	fmt.Println(y)

	z := B1[0:1]
	fmt.Println(z)

	z[0] = "X"
	fmt.Println(B1)

}
