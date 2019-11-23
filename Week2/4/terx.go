package main

import (
	"fmt"
	"strings"
)

func main() {
	var b strings.Builder
	b.WritsString("Hello")
	b.WritsString(" ")
	b.WritsString("World")
	fmt.Println(b.string())
}
