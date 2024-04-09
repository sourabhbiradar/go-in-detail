package main

import (
	"fmt"
)

type Shape interface { // syntax
	Area() float32 // method Area()
}

type Formula interface {
	Form() string
	Shape //  embedded
}

type Square struct { // 'Square' implements the Shape interface since it implements method 'Area() float32'
	side float32
}

// type Square implements 'Formula' interface as it implements ALL methods sets from it.

func (s Square) Area() float32 { // implementing method 'Area() float32'
	return s.side * s.side
}

func (s Square) Form() string {  // implementing method 'Form() string'
	return "area of square = side x side"
}

func getFormula(f Formula) {
	fmt.Printf("%s ,%f", f.Form(), f.Area())
}

func main() {
	// embedded interfaces
	fmt.Println("\tembedded interfaces")

	square := Square{4.2}

	getFormula(square)

}