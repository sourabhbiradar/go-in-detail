package main

import (
	"fmt"
)

func main() {
	// iota
	fmt.Println("\tiota")

	const x = iota // x = 0
	fmt.Println("default iota value :", x)

	const ( // // only be used with const
		C0 = iota
		C1 = iota
		C2 = iota
	)
	fmt.Println("C0, C1, C2 :", C0, C1, C2) // sequence of numbers frm zero

	const (
		i0 = iota - 1 //0 - 1 = -1
		i1
		_ // skips i2 == 1
		i3
		i4
		i5
	)
	fmt.Println("starting value frm '-1' :", i0, i1, i3, i4, i5)

	const (
		a = iota + 12 //0 + 12 = 12
		b
		c
		d
	)
	fmt.Println("starting value frm '12' :", a, b, c, d)

	const (
		i = iota * 3 // 0 x 3 = 0
		j
		k
	)
	fmt.Println("multiples of '3' :", i, j, k)

}