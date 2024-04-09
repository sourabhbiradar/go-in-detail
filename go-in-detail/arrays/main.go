package main

import "fmt"

func main() {
	// Ways to declare Arrays

	// way 1 to initialize and declare an array
	var arr [5]int = [5]int{10, 9, 8, 7, 6}
	//  array 'arr' of length 5
    // once length is declared it can't be changed
	fmt.Println("arr :", arr)

	// way 2
	var arr1 [4]int
	arr1[0] = 11 // index 0 --> 11
	arr1[1] = 12 // index 1 --> 12
	arr1[2] = 13 // index 2 --> 13
	arr1[3] = 14 // index 3 --> 14
	fmt.Println("arr1 :", arr1)

	//way 3
	var arr2 = [3]string{"a", "b", "c"}
	fmt.Println("arr2 :", arr2)

	// way 4
	arr3 := [3]float32{1.01, 2.02, 3.03}
	fmt.Println("arr3 :", arr3)

	// way 5 --> sparse array
	arr4 := [5]int{5, 2: 15, 20}
	// 2 : 15 --> 15 is in index 2
	fmt.Println("arr4 :", arr4)
	// since index 1 and 5 are not entered they will assigned default value of zero '0'

	// way 6
	arr5 := [...]string{"w", "x", "y", "z"}
	// ellipces [...] --> based on elements length of array is determined
	fmt.Println("arr5 :", arr5)
    
}