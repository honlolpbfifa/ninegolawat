package main

import "fmt"

func main() {
	defer fmt.Println("Hello world")
	var a map[int]int
	a[0] = 1
}
