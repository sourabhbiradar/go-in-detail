package main

import "fmt"

func main(){
	// Primitive or Basic Data type
	// string , numbers and boolean
	
	// string
	var strVal string
	fmt.Println(strVal) 	// default value --> empty string " "
	strVal = "Hey" 			// value inside double quotes ""
	fmt.Println(strVal)

	// numbers --> int ( integer )
	var intVal string
	fmt.Println() 			// default value --> 0 (zero)
	intVal = 38 			// no double quotes
	fmt.Println(intVal)

	// boolean ( bool ) --> true or false
	var boolVal string
	fmt.Println() 			// default value --> false
	boolVal = true      	// no double quotes
	fmt.Println(boolVal)
}