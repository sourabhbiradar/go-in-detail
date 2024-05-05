package main

import (
	"fmt"
	"time"
)

//func main() {
	var counter int

	for range 10 { // go 1.22 onwards

		go func() { // initializing go routine with `go`keyword
			counter++

		}() // calling go routine
	}
	time.Sleep(3 * time.Second) // without sleeping counter may NOT be inceremented 10 times becoz we wud exit program before even all go routines got executed.
	fmt.Println("Expected:10 ,Got:", counter)
}
