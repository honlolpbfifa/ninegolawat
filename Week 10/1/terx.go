package main

import (
	"fmt"
	"io"
)

func main() {
	reader := string.NewReader("HelloWorld")
	P := make([]byte, 3)
	for {
		n, err := reader.Read(p)
		if err == io.EOF {
			break
		}
		fmt.Println(string(p[:n]))
	}
}
