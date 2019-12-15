package main

import "fmt"

func suntract(number int) {
	*number = *number - 2
}

func main() {
	x := 10
	suntract(&x)
	fmt.Println(x)

}
