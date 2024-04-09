package main

import (
	"fmt"
)

func printVal(a interface{}) {   // param 'a' is of type 'any' 
	fmt.Printf("'%v' of type %T\n", a, a)
}

func main() {
	// empty interfaces
	fmt.Println("\tempty interfaces")

	anyType := []interface{}{1, "one ", 1.01, true}

	for _, val := range anyType {
        printVal(val)
    }
}