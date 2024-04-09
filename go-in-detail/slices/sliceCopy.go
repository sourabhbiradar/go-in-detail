package main

import "fmt"

func main() {
	// Slices

	// coping slice to slice a.k.a. slicing

	src := []string{"s", "l", "i", "c", "e"}
	fmt.Println("src :", src)

	slc := src[2:]
	// copy elements index[2] to last element to 'slc' from 'src'

	fmt.Println("child slice :", slc)

	// manuplilating child slice 'slc' alters parent slice 'src' and vice versa
	// this is called ripple effect
	src[3] = "d" // src[3] --> "c" to "d"
	fmt.Println("altered src :", src)
	fmt.Println("altered child slice slc:", slc)

	// coping usin copy() function
	a := []int{10, 20, 30, 40}
	b := make([]int, 5)

	copy(b, a)              // copy(to,from)
	fmt.Println(b)          // [10 20 30 40 0]
	fmt.Println(copy(b, a)) // how many elements are copied --> 4

	// coping from array
	arr := [5]string{"a", "r", "r", "a", "y"}
	slice := make([]string, 5)

	copy(slice, arr[:])  // index [0] to index[4]
    // specified portion of arr can also be copied, try --> arr[1:3]

	fmt.Println("Copied from arr :", "slice :", slice)

}