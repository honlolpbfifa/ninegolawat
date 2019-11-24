package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.trim("-Hello - World-", "-"))
	fmt.Println(strings.Trim("+Hello World-", "-+"))

}
