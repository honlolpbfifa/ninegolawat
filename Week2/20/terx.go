package main

import "fmt"

const greeting = "Hello World"

func main() {
	fmt.Println(greeting)
	greeting = "Hello World"
	fmt.Println(greeting)

}
