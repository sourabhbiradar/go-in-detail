package main

import (
	"fmt"
)

func clsr() func() int {
	i := 0  // replace i -> any int
	return func() int {
		i++
		return i 
	}
}

func main() {
	// function closures
	fmt.Println("\tclosures")

	newClsr := clsr()
	fmt.Println(newClsr())
	fmt.Println(newClsr())
	fmt.Println(newClsr())     // everytime its called value increases
	//fmt.Println(newClsr())

	newClsr2 := clsr()
	fmt.Println(newClsr2())     // again from 1 

}