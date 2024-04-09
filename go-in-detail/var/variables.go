package main

import "fmt"

var globalVar string = "global variables can be used anywhere in program" // cant use short hand for global variables

func main() {
	// var --> variables
	// these are different ways how variables can be declared n intialized

	var a string = "hello" // explicit
	fmt.Println(a)

	var b = "bey" // implicit as value is string 'b' hence its of type string
	fmt.Println(b)

	c := "fine" // short hand declaration n intialization
	fmt.Println(c)

	var i, j int // multiple variables declared in same line of same type
	i = 3
	j = 4
	fmt.Println(i, j)

	var (
		x float32 = 4.22
		y string  = "using braces ()"
		t bool    = !true
	)
	// multiple variables of different types are grouped with in braces ( )
	fmt.Println("x :", x, "y :", y, "t :", t)

	var x1, y1, f1 = 2, "using implicit", false // multiple variables of different types
	fmt.Println("x1 :", x1, "y1 :", y1, "f1 :", f1)

	x2, y2, f32 := 10, "using short hand", 3.14 // multiple variables of different types
	fmt.Println("x :", x2, "y2 :", y2, "f32 :", f32)

	// above variables are all local or function variables so, cant be used outside function block.

	fmt.Println(globalVar) // globalVar

}