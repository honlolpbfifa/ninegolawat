package main

import "fmt"

func sum(c chan int, number ...int) {
	sum := 0
	for _, V := range number {
		sum = sum + V
	}
	c <- sum
}
func printer(c chan int) {
	number := <-c
	fmt.Println(number)
}
func main() {
	c := make(chan int)
	go printer(c)
	go printer(c)
	go sum(c, 1, 2, 3)
	go sum(c, 10, 11)

	var input string
	fmt.Scan(&input)

}
