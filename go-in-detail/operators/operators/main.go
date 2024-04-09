package main

import "fmt"

func main() {
	// operators

	const num int = 2 // assignment

	// arithmatic operators
	fmt.Println("\t arithmatic operators")
	var a, b int = 10, 5
	fmt.Println(a + b) // addition
	fmt.Println(a - b) // substraction
	fmt.Println(a * b) // multiplication
	fmt.Println(a / b) // division
	fmt.Println(a % b) // reminder

	// comparsion operators
	fmt.Println("\t comparsion operators")
	fmt.Println(4 == 4)         // equal to
	fmt.Println(4 != 5, 4 != 4) // NOT equal to
	fmt.Println(4 < 5)          // lesser than
	fmt.Println(3 > 2)          // greater than
	fmt.Println(4 <= 5)         // lesser than or equal to
	fmt.Println(5 >= 5)         // greater than or equal to

	// logical operators
	fmt.Println("\t logical operators")
	fmt.Println(10 > 5 && 4 < 5) // && --> AND
	fmt.Println(10 < 5 || 5 > 4) // || --> OR
	fmt.Println(!true)           // ! --> NOT

	// assignment operators
	fmt.Println("\t assignment operators")
	a += 2         // addition assignment
	fmt.Println(a) // a += 2 --> a = a + 2
	b -= 2         // substraction assignment
	fmt.Println(b) // a -= 2 --> a = a - 2

	var x, y = 10, 20
	x++              // increment operator
	fmt.Println(x)   // x++ --> x + 1
	y--              // decrement operator
	fmt.Println(y)   // y-- --> y - 1

}