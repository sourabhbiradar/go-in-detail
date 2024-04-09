package main

import (
	"fmt"
)

type Shape interface { // syntax
	Area() float32 // method Area()
}

type Square struct { // 'Square' implements the Shape interface since it implements method 'Area() float32'
	side float32
}

type Circle struct { // 'Circle' implements the Shape interface since it implements method 'Area() float32'
	radius float32
}

func (s Square) Area() float32 { // implementing method 'Area() float32'
	return s.side * s.side
}

func (c Circle) Area() float32 { // implementing method 'Area() float32'
	return 3.14 * c.radius * c.radius
}

func getArea(s Shape) {
	fmt.Println("getArea()")
	fmt.Printf("area :%f\n", s.Area()) // since Shape has Area()
}

func main() {
	// interfaces
	fmt.Println("\tinterfaces")

	square := Square{4.2}
	circle := Circle{radius: 5.0}

	getArea(square)
	getArea(circle)

}