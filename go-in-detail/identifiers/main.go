package main

import "fmt"

func pink() {
	fmt.Println("'pink' here is identifier")
}

func main() {
	// identifiers

	var newInt int = 76 // newInt is known as idetifier
	fmt.Println("newInt :", newInt)

	var _next string = "_next" // identifiers can either start with letter or underscore
	fmt.Println("_next :", _next)

	var ident3 bool = true // valid identifier
	fmt.Println("ident3 :", ident3)

	var bool int = 5 // its valid but NOT recommanded to use keywords
	fmt.Println(bool)

	pink()     // func pink()
}