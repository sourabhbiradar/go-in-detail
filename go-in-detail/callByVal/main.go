package main

import (
	"fmt"
)

func main() {
	// call by value
	fmt.Println("\tcall by value")

	a := 14
	fmt.Println("before a:", a)
	funcA(a)
	fmt.Println("after a:", a) // updating 'x' will NOT effect actual 'a' value

	// call by reference
	fmt.Println("\tcall by refernce")

	// passed value as pointer
	fmt.Println("passed value as pointer")
	var b *int
	b = &a
	fmt.Println("before a:", a)
	funcB(b)
	fmt.Println("after a:", a) // updating 'x' WILL effect actual 'a' value

	// passed value as slice
	fmt.Println("passed value as slice")
	var s []int = []int{1, 2, 3}
	fmt.Println("before s:", s)
	funcC(s)
	fmt.Println("after s:", s) // appending in funcC wont effect

	// passed value as map
	fmt.Println("passed value as map")
	var m map[int]string = map[int]string{1: "one", 2: "two", 3: "three"}
	fmt.Println("before m:", m)
	funcD(m)
	fmt.Println("after m:", m) // m[2] deleted & m[6]="six" is appended
}

// call by value
func funcA(x int) { // x = copy of 'a'
	x = 28
}

// call by ref
// passed value as pointer
func funcB(x *int) {
	*x = 28
}

// passed value as slice
func funcC(x []int) {
	x[0] = 8         // == updating s[0] to 8
	x = append(x, 5) // appending has NO effect on 's'
}

// passed value as map
func funcD(x map[int]string) {
	delete(x, 2)
	x[6] = "six"
}