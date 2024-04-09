package main

import "fmt"

func main() {
	// float data type

	var smallFloat float32
	fmt.Println(smallFloat) // default value --> 0
	smallFloat = 4.0123459
	fmt.Println(smallFloat) // 4.012346

	var bigFloat float64 = 5.0123456789101119
	fmt.Println(bigFloat) 	// 5.012345678910112

}