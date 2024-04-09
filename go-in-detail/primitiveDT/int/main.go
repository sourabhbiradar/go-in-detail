package main 

import "fmt"

func main(){
	// int Data Type
	
	var intVal int
	fmt.Println(intVal)         // default value is 0 (zero)
	
	intVal = -37
	fmt.Println(intVal) 		// can be negetive
	
	intVal = 23.76
	fmt.Println(intVal)         // drops value after decimal point --> 23 output
    
	intVal = 3
	fmt.Println(intVal+4)       // 7  ( 3+4 )

	intVal = 5
	fmt.Println(intVal*intVal)  // 25  ( 5x5 )

	val := 33.09
	fmt.Println(int(val))       // 33

	// int(value) can also be used to convert other number data type or value to integer known as 'type casting'

	// if u r using specific int types int8 , uint32 etc. make sure to use values within their range.

	var runeVal rune := 'A'     
	fmt.Println(runeVal)  		// 65 ASCII value of 'A'

	var byteVal byte := 'a' 	
	fmt.Println(byteVal)        // 97 ASCII value of 'a'
	
	// please read README.md

}