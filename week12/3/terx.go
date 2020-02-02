package main

import (
	"fmt"
	"time"
)

func say(txt string) {
	for i := 0; i < 3; i++ {
		fmt.Println(time.Now(), ":", i, ":", txt)
		time.Sleep(time.Millisecond)
	}
}
func main() {
	go say("Hello")
	go say("Hi")
}
