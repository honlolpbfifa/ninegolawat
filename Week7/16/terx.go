package main

import "fmt"

func main() {
	fmt.Println(567888)

	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	defer fmt.Println(4)

}
