package main

import "fmt"

func main() {
	// Arrays
	// two dimensional --> 2D array

	arr2D := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("2D array :", arr2D)
	// [2][3]int --> [2] arrays each having size of [3]

    // accessing single element from arr2D
    fmt.Println("arr2D[1][2] :", arr2D[1][2]) // 6
    // from array at index [1] element from index[2]

	//accessing sub arrays of arr2D
	fmt.Println("arr2D[0] :", arr2D[0]) // first array of arr2D
	fmt.Println("arr2D[1] :", arr2D[1]) // second array of arr2D

	// every element from arr2D
	fmt.Println("arr2D[0][0] :", arr2D[0][0])
	fmt.Println("arr2D[0][1] :", arr2D[0][1])
	fmt.Println("arr2D[0][2] :", arr2D[0][2])
	fmt.Println("arr2D[1][0] :", arr2D[1][0])
	fmt.Println("arr2D[1][1] :", arr2D[1][1])
	fmt.Println("arr2D[1][2] :", arr2D[1][2])

}