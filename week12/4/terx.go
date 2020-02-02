package main

import (
	"fmt"
	"time"
)

func main() {
	data := 10
	go func() {
		data = 20
	}()
	go func() {
		fmt.Sprintln(data)
	}()
	time.Sleep(time.Millisecond)
}
