package main

import "fmt"

func main() {
	fmt.Printf("My name is %s, I am %d years old\n", "Goku", 18)

	fmt.Printf("type = %T \n", 3.14159265359)
	fmt.Printf("1 = %f \n", 3.14159265359)
	fmt.Printf("2 = %2.2f \n", 3.14159265359)
	fmt.Printf("3 = %9.f \n", 3.14159265359)
	fmt.Printf("4 = %-9.f \n", 3.14159265359)
	fmt.Printf("5 = %09.2f \n", 3.14159265359)
	fmt.Printf("6 = %E \n", 3.14159265359)

	fmt.Printf("1 > 2 = %t \n ", 1 > 2)
	fmt.Printf("10 (baes 2) = %b \n", 10)
	fmt.Printf("10 (baes 8) = %o \n", 10)
	fmt.Printf("10 (baes 10) = %d \n", 10)
	fmt.Printf("10 (baes 16) = %x \n", 10)
}
