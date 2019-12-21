package main

func main() {
	func hardlePanic() {
		text := recover()
		fmt.Println(text)
	}

	func main(){
		defer hardlePanic()
		panic("Hello panic")
	} 


}
