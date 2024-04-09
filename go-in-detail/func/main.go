package main

import "fmt"

func main() {
	// functions
	// func
	fmt.Println("\tfunc --> function")

	// simple function
	function() // calling function()
	// function needs to be called in order to be executed

	// func that takes input
	input(3) // passing parameter/value to 'func input()'

	// func that take multiple inputs and returns value
	fmt.Println("func returnVal()")
	fmt.Println(" res =", returnVal(3, 4)) // values for 'n' and 'm'

	// multiple return values
	fmt.Println("func multiReturn()")
	name, jersey := multiReturn(18, "Virat")  // each variable catches one  return value
	fmt.Println(name, jersey)

}

func function() {
	fmt.Println("func function()")
}

func input(val int) { // takes input 'val'
	sum := val + val
	fmt.Println("func input()\n sum =", sum)
}

func returnVal(n, m int) int { // takes 2 values 'n' 'm'
	res := n + m
	return res // returns 'res' of type 'int'
}

// func returnVal(n, m int) (res int) { // takes 2 values 'n' 'm' & returns 'res' of type 'int'
// 	res = n + m
// 	return // just return
// }

func multiReturn(a int, b string) (string, int) {
	return b, a
}