package main

import (
	"fmt"
)

// type person struct{
//    name string
//    age int
//    city string
// }

type person struct {
	name, city string // as name & city are of same type
	age        int
}

func main() {
	// ways to initialize & declare structs
	fmt.Println("\tways to initialize & declare structs")

	// way 1
	p := person{name: "Abc", age: 20, city: "city"} // with fields
	fmt.Println("p :", p)
	fmt.Printf("p :%+v \n", p) // '%+v' prints values with fields

	// way 2
	p1 := person{"Xyz", "City", 20} // WITHOUT fields n should be in sequence as declared
	fmt.Println("p1 :", p1)

	// way 3
	NewP2()

	// way 4
	fmt.Println("p3 :", NewP3())

	// way 5
	fmt.Println("p4 :", NewP4())

	//way 6
	fmt.Println("p5 :", *NewP5())
}

// way 3
func NewP2() {
	var p2 person = person{
		name: "Aaa",
		city: "c",
		age:  20,
	}
	fmt.Println("p2 :", p2)
}

// way 4
func NewP3() person { // returning 'person' with values
	p3 := person{"Ccc", "c", 20}
	return p3
}

// way 5
func NewP4() person {
	var p4 person
	p4.name = "Bbb" // instance with field names
	p4.city = "c"
	p4.age = 20
	return p4
}

// way 6
func NewP5() *person {
	var p5 = new(person) // using 'new' keyword
	p5.name = "Bbb"      // instance with field names
	p5.city = "c"
	p5.age = 20

	fmt.Println("NewP5() p5 :", *p5) // dereferncing p5
	return p5
}
