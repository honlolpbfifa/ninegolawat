package main

import "fmt"

func sun(numbers ...int) int {
	total := 0
	for _, n := range numbers {
		total = total + n
	}
	return total
}
func main() {
	x := sun(1, 3, 5, 7, 9)
	fmt.Println

}
