package main

import "fmt"

func main() {
	// bits and bitwise operators

	var a, b uint8 = 0, 1
	// always use 'uint' while workin with bits

	// OR --> '|'
	fmt.Println("OR operator")
	c := a | b
	fmt.Printf("%v | %v --> %v\n", a, b, c)
	c = a | a
	fmt.Printf("%v | %v --> %v\n", a, a, c)
	c = b | a
	fmt.Printf("%v | %v --> %v\n", b, a, c)
	c = b | b
	fmt.Printf("%v | %v --> %v\n", b, b, c)

	// AND --> '&'
	fmt.Println("AND operator")
	d := a & b
	fmt.Printf("%v & %v --> %v\n", a, b, d)
	d = a & a
	fmt.Printf("%v & %v --> %v\n", a, a, d)
	d = b & a
	fmt.Printf("%v & %v --> %v\n", b, a, d)
	d = b & b
	fmt.Printf("%v & %v --> %v\n", b, b, d)

	// XOR --> '^'
	fmt.Println("XOR operator")
	e := a ^ b
	fmt.Printf("%v ^ %v --> %v\n", a, b, e)
	e = a ^ a
	fmt.Printf("%v ^ %v --> %v\n", a, a, e)
	e = b ^ a
	fmt.Printf("%v ^ %v --> %v\n", b, a, e)
	e = b ^ b
	fmt.Printf("%v ^ %v --> %v\n", b, b, e)

}