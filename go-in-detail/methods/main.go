package main

import (
	"fmt"
)

type Rectangle struct {
	height int
	base   int
}

func (r Rectangle) AreaR() int { // method func that works with 'person' struct , 'r' is called value reciever.
	area := r.base * r.height
	return area
}

func (r *Rectangle) ChangeH(h int) { // 'r' pointer reciever
	r.height = h
}

func main() {
	// methods
	fmt.Println("\tmethods")

	// struct method
	fmt.Println("struct method")

	// value reciever

	fmt.Println("value reciever")
	rect := Rectangle{height: 5, base: 6}
	// methods are called through the type they are associated with.

	x := rect.AreaR()
	fmt.Println("AreaR() :", x)

	// pointer reciever
	fmt.Println("pointer reciever")
	rect.ChangeH(3)
	fmt.Println("ChangeH() :", x)

}