package main

import "fmt"

func main() {
	// string data type

	var str string = "hello"
	fmt.Println(str)

	str = "hello\nworld" // '\n' new line
	fmt.Println(str)     // hello and world are prited in differnt lines

	str = `a b.  c      d
    x y z`
	// back ticks (not single quotes)
	fmt.Println(str) // prints as raw data

	var firstName, lastName string
	firstName = "Abc"
	lastName = "Xyz"

	fmt.Println(firstName + " " + lastName) // concatination
	// output --> Abc Xyz

}