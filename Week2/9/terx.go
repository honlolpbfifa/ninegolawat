package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.HasPrefix("Hello World", "hello"))
	fmt.Printf(strings.HasPrefix("Hello World", "Hello"))

}
