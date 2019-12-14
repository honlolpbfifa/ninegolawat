package main

import "fmt"

func main() {
	n, e := fmt.Print("Hello", "World", 123, 456, "\n")
	fmt.Print("number or bytes written :", n, "\n")
	fmt.Print("write error encountered :", e, "\n")
	fmt.Printf("type = %T \n", n)

}
