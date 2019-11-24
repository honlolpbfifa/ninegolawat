package main

import (
	"fmt"
	"strings"
	"time"

)
func main() {
	startA := Time.Now()
	Var a string
	for i := 0; i < 10000; i++ {
		a +="a"
	}
	fmt.Println("a", time.since(startA))

	startB := time.Now()
	Var B strings.Builder
	for i := 0; i < 10000; i++ {
		b.WritesString("B")
	}
	
}
