package main

import "fmt"

func main() {
	// Arrays

	a := [3]string{"a", "b", "c"}
	// len()
	fmt.Println("length of arr is ", len(a)) // len(array_name)

	// accessing elements using index
	var arr [5]int = [5]int{9, 8, 7, 6, 5}

	fmt.Println("first element is at arr[0] =", arr[0]) // arr[index]
	fmt.Println("second element is at arr[1] =", arr[1])
	fmt.Println("third element is at arr[2] =", arr[2])
	fmt.Println("forth element is at arr[3] =", arr[3])
	fmt.Println("fifth element is at arr[4] =", arr[4])

	fmt.Println(arr[3])                            // 6
	fmt.Println("last element :", arr[len(arr)-1]) // [len()-1]
	//  fmt.Println(arr[6])    // error : out of bound

	//update element of array
	arr[4] = 10        // arr[4] was 5 but now its reassigned as 10
	fmt.Println("updated arr[4] :", arr[4])  

	var b = [5]string{"a", "b", "c", "d", "e"}

	fmt.Println("b[1:4] :", b[1:4])
	// elements from index[1] to index[3]
	// index[4] not included

	fmt.Println("b[:4] :", b[:4])
	// index[0] to index[3]

	fmt.Println("b[2:] :", b[2:])
	// index[2] to last element

}