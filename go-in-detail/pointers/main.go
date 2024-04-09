package main

import (
	"fmt"
)

// pointer to struct
type info struct {
	name string
	age  int
}

func main() {
	// pointers
	fmt.Println("\tpointers")

	var x *int
	fmt.Println(x) // <nil> since no value is stored

	fmt.Println("pointerToInt")
	var pointerToInt *int // pointerToInt holds the address of 'int' variable
	age := 10
	pointerToInt = &age     // holds the address of 'a' (int variable)
	ageVal := *pointerToInt // dereferencing pointerToInt

	fmt.Println(age, pointerToInt, ageVal)

	// working with copy
	fmt.Println("\tworking with copy")
	a := 4

	//increment(a) // uncomment this line
	fmt.Println(a) // we worked with copy of 'a' so no change in actual value of 'a'

	increment(&a)  // directly worked with value of 'a'
	fmt.Println(a) // hence the change

	// pointer to struct
	fmt.Println("\tpointer to struct")
	p := info{
		"Abc",
		18,
	}
	fmt.Println("before updatin p :", p)
	person(&p)
	fmt.Println("updated p ;", p)

}

// working with copy
func increment(i *int) { //  i = copy of 'a'
	*i++
}

// pointer to struct
func person(i *info) {
	i.name = "Dr.Abc" // i.name == *i.name ,so do NOT derefence i.name or i.age
	i.age = 23
	fmt.Println(i)
}