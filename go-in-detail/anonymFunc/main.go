package main

import "fmt"

func main() {
	// anonymous func
	fmt.Println("\anonymous func")

	func(str string) {
		fmt.Println(str)
	}("func ()") // passing arguements

	fn := func(str string) {
		fmt.Println(str)
	}

	fn("'fn' is also anonymous func")

}