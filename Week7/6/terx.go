package main

import "fmt"

func writeline(a ...interface{}) {
	for _, v := range a {
		fmt.Println(v)
	}
}
func main() {

}
