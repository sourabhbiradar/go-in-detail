package main

import "fmt"

func main() {
	// Slices

	// slices are also declared and initilized like arrays.

	// Can ALSO be declared using make() function

	slc := make([]string, 5)
	// slice of type string and size 5

	slc[0] = "s"
	slc[1] = "l"
	slc[2] = "i"
	slc[3] = "c"
	slc[4] = "e"

	fmt.Println("slc :", slc)

	// using size and capacity
	slc2 := make([]int, 3, 6)
	// no need to specify capacity Go will do it by itself

	slc2[0] = 1
	slc2[1] = 2
	slc2[2] = 3

	fmt.Println("slc2 :", slc2)

	fmt.Println("size of slc2 :", len(slc2))     // len()
	fmt.Println("capacity of slc2 :", cap(slc2)) // cap()

	// append()
	slc2 = append(slc2, 10) //new element --> 10
	// x = append(x,element)
	fmt.Println("new element added to slice :", slc2)

	fmt.Println("updated size of slc2 :", len(slc2)) // new size = 4

	// append multiple elements
	slc2 = append(slc2, 20, 30, 40) //new element --> 20 ,30,40

	fmt.Println("new elements added to slice :", slc2)

	fmt.Println("updated size of slc2 :", len(slc2)) // new size = 7
	fmt.Println("capacity of slc2 :", cap(slc2))     // new capacity = 12

	// appending slice to another slice
	slc3 := []int{-99, -88, -77}
	slc4 := []int{11, 22, 33}

	slc4 = append(slc4, slc3...) // slc3... --> slice is broken into single single elements

	fmt.Println("appended slc3 to slc4 :", slc4)

}