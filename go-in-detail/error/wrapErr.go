package main

import (
	"errors"
	"fmt"
)

func funcOne() error {
	return fmt.Errorf("Error in funcOne()")
}

func funcTwo() error {
	firstErr := funcOne()
	if firstErr != nil {
		return fmt.Errorf("Error at funcTwo : %w", firstErr) // wrapping funcOne() error [firstErr] in funcTwo() error
		// '%w' for holding wrapped error.
	}
	return nil
}

func main() {
	// wrapping & unwrapping errors
	fmt.Println("\twrapping & unwrapping errors")

	// wrapping error
	fmt.Println("wrapping errors")

	f2 := funcTwo()
	fmt.Println(f2)

	// unwrapping error
	fmt.Println("unwrapping errors")

	fTwo := funcTwo()
	unwErr := errors.Unwrap(fTwo)   // errors.Unwrap()
	fmt.Println("original error :", unwErr) // after unwrapping

}