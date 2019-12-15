package main

import "fmt"

func say(greet string) func(string) func(int)string int {
	return func(name string) string {
		return func(age int) int {
			return greet + name + age
		}
	}
}

func main() {
	x := say("Hello")

	fmt.Println(x("goku"))
	fmt.Println(x("gohan"))
	fmt.Println(x(18))

}
