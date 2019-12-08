package main

import "fmt"

func main() {
	B1 := [6]string{"A", "B", "C", "D", "E", "F"}
	fmt.Println(B1)

	x := B1[0:2]
	fmt.Println(x)

	y := x[2:4]
	fmt.Println(y)

	z := y[0:4]
	fmt.Println(z)

	z[0] = "X"
	fmt.Println(B1)

}
