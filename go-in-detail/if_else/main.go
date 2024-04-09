package main

import "fmt"

func main() {
	// conditional statements
	// if , if-else , if-else-if

	// if statement
	fmt.Println(" 'if' statement ")

	if true {
		fmt.Println("condition is true so code within 'if' is executed")
	}

	colour := "blue" // try "red"
	if colour == "blue" {
		fmt.Println(" coloue == blue so 'if' block is executed ")
	}
	// colour := "red" --> condition (colour == blue) is false so 'if' block will NOT be executed.
	// colour := "blue" --> condition true so  print statement will be printed.

	// if-else statement
	fmt.Println(" 'if-else' statement ")

	var age = 20 // try 16
	if age >= 18 {
		fmt.Println("you are an adult") // if condition is true --> executed
	} else {
		fmt.Println("you are NOT an adult") // if condition is false --> executed
	}

	// if-else-if statement
	fmt.Println(" 'if-else-if' statement ")

	age2 := 15
	if age2 >= 18 { // false
		fmt.Println("you are an adult") // NOT executed
	} else if age2 >= 13 && age2 <= 17 { // true
		fmt.Println("you are a teenager") // executed
	} else if age2 >= 10 && age2 <= 12 { // // addtional 'else-if'
		fmt.Println("you are pre-teen")
	} else {
		fmt.Println("you are a child") // NOT executed
	}

	// initializing varibale in 'if'
	fmt.Println("initializing varibale in 'if'")

	if n := isEven(age); n { // line 26 --> var age = 20
		fmt.Println("age is even ")
	}

	isEven(8)

}

func isEven(num int) bool { // you will learn func or functions in later codes
	return num&1 == 0
}