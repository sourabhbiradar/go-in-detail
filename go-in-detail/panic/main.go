package main

import (
	"fmt"
)

func fix() {
	r := recover() // recover()
	if r != nil {
		fmt.Println("Recovered from panic")
	}
}

func panicFunc() {
	//	defer fix()      // uncomment this line to RECOVER FROM PANIC
	panic("oops!! something went wrong")
}

func example() {
	fmt.Println("deferred func example()")
}

func main() {
	// panic & recovery
	fmt.Println("\tpanic & recovery")

	defer example() // will be Executed

	defer func() {
		fmt.Println("func main() deferred") // will be Executed
	}()

	fmt.Println("start")

	panicFunc()

	fmt.Println("exit") // will NOT be executed when PANICED bt executed after recovery

}