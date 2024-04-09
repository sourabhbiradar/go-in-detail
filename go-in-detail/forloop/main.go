package main

import "fmt"

func main() {
	// for loop

	// for loop with a condition
	fmt.Println("\tfor loop with a condition")
	a := 0
	for a < 4 {
		fmt.Println("a :", a) // executes until i < 4
		a++                   // i= i+1 --> increase i by 1
	}

	// infinite for loop
	// break
	fmt.Println("\tinfinite for loop")

	for true {
		fmt.Println("loop runs for ever as condition is true") // comment 'break' in below line

		break // to break out loop we use 'break'
	}

	// for loop with for clause
	fmt.Println("\tfor loop with for clause")
	for i := 1; i < 6; i++ { // initialization ; condition ; post statement
		fmt.Println("i :", i)
	}

	// for loop with range cluase
	fmt.Println("\tfor loop with range cluase")
	arr := [4]int{10, 20, 30, 40}

	for index, elem := range arr { // iterates over every element from array 'arr'
		fmt.Printf("index : %v , element : %v\n", index, elem)
	}

	// iterating over string

	str := "coder"
	for _, letter := range str { // you can omit any iteration varible by replacing variable with '_'
		fmt.Printf("letter : %c\n", letter) // %c --> string(letter)
	}

	// iterating over map
	m := map[int]string{1: "one", 2: "two", 3: "three"}
	for key, val := range m {
		fmt.Printf("key %d : value %s\n", key, val)
	}

}