package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Replace("Hello World", "1", "x", 2))
	fmt.Println(strings.Replace("Hello World", "1", "x", -1))

}
