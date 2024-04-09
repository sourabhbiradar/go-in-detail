package main

import (
	"fmt"
)

func main() {
	// defer
	fmt.Println("\tdefer")

	defer one() // executed at last
	defer two() // two () is executed before one() --> LIFO
	three()     // even though three() is called last,it is executed before one() and two()
	// panic("oops!!")  // defer are not executed

}

func one() {
	fmt.Println("func one()")
}

func two() {
	fmt.Println("func two()")
}

func three() {
	fmt.Println("func three()")
}