package main

import (
	"errors"
	"fmt"
)

var ErrFirst = CustomErr{"first error"} // custom sentinel error , starts with Err....
// swap CustomErr{"first error",400}

type CustomErr struct {
	message string
	// status  int //  can be accessed using errors.As()
}

func (c CustomErr) Error() string { // func() Error() always
	return c.message
}

func firstErr() error {
	return fmt.Errorf("Error at firstErr() : %w", ErrFirst)
	// return ErrFirst
}

func main() {
	// errors.Is() & errors.As()
	fmt.Println("\terrors.Is() & errors.As()")

	// errors.Is()
	fmt.Println("errors.Is()")

	err := firstErr()
	if errors.Is(err, ErrFirst) { // comapring error from firstErr() with sentinel ErrFirst
		fmt.Println("sentinel error found")
	}
	fmt.Println(err)

	// error..As()
	fmt.Println("error.As()")

	var custErr CustomErr
	if errors.As(err, &custErr) {
		fmt.Println("extracted customErr :", custErr.message)  // custErr.status
	}

}