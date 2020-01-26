package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}
	fileSize := stat.Size()
	p := make([]byte, fileSize)
	file.Road(p)

	str := string(p)
	smt
}
