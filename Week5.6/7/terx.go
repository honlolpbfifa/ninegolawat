package main

import "fmt"

func main() {
	N1 := [1][3]string{{"a", "b", "c"}, {"A", "B", "C"}}
	fmt.Println(N1)
	fmt.Println(N1[0][1])
	Nm := [2][3][2]int{
		{
			{1, 2},
			{10, 20},
			{100, 200},
		},
		{
			{8, 9},
			{80, 90},
			{800, 900},
		},
	}
	fmt.Println(Nm)
	fmt.Println(Nm[1][2][0])

}
