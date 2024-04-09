package main

import "fmt"

var global string = "global variable"

func funcBlock() {
	var funcVar = "function variable"
	fmt.Println("accessible from funcBlock()", global)
	fmt.Println("funcVar :", funcVar) // funcVar can only be accessed in this function
}

func main() {
	// code block and shadowing

	// code block and scoping

	// package-level scope
	fmt.Println("\tpackage-level scope")
	fmt.Println("accessible from anywhere in func main()", global)
	funcBlock()

	// function-level
	fmt.Println("\tfunction-level scope")
	// uncomment below line
	// fmt.Println("funcVar :", funcVar)   // funcVar can NOT be accessed outside funcBlock()

	// statement-level
	fmt.Println("\tstatement-level scope")
	if true {
		stateVar := "statement variable"
		fmt.Println("stateVar", stateVar)
	}
	// uncomment below line
	// fmt.Println("stateVar", stateVar)   // statementVar can NOT be accessed outside 'if{}'

	// block litral or anonymous block
	fmt.Println("\tanonymous block scope")
	{
		var anonymVar = "anonymous block variable"
		fmt.Println("anonymVar :", anonymVar)
	}
	// uncomment below line
	// fmt.Println("anonymVar", anonymVar)   // anonymVar can NOT be accessed outside '{}'

	// inner and outer block
	fmt.Println("\tinner and outer block")
	if true {
		outerVar := "outer variable"
		fmt.Println(" accessing outerVar from outer block :", outerVar)
		// uncomment below line
		// fmt.Println("accessing innerVar from outer block :",innerVar)   // can NOT
		if true {
			innerVar := "inner variable"
			fmt.Println(" accessing innerVar from inner block :", innerVar)
			fmt.Println(" accessing outerVar from inner block :", outerVar)
		}
	}

	// shadowing
	fmt.Println("\tshadowing")

	var colour = "Red"
	fmt.Println("outer block ", colour)
	if true {
		var colour = "Blue"
		fmt.Println("inner block ", colour) // inner variable given prioroty over outer variable
	}
	fmt.Println("after 'if' outer block ", colour) // since inner variable is NOT accessible so colour == "Red"

}