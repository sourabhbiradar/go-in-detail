package main

import (
	"errors"
	"fmt"
)

func calculate(x, y float32) (float32, error) {
	if y == 0 {
		return 0, errors.New("\nCan NOT divide by zero") // errors.New()
	}
	return x / y, nil
}

func divide(x, y float32) (float32, error) {
	if y == 0 {
		return 0, fmt.Errorf("can't divide by %f", y)
	}
	return x / y, nil
}

// using error interface
type CustomErr struct {
	code    int
	message string
}

func (c CustomErr) Error() string {
	return fmt.Sprintf("Error : %d ,%s", c.code, c.message)
}
func solution(x, y float32) (float32, error) {
	if y == 0 {
		return 0, CustomErr{code: 1, message: "can NOT divide by zero"}
	}
	return x / y, nil
}

func main() {
	// errors & error handling
	fmt.Println("\terrors & error handling")

	// error handling using errors.New()
	fmt.Println("errors.New()")

	c, err := calculate(3, 0)
	if err != nil {
		fmt.Println("Error :", err)
	}
	fmt.Println(c)

	// using fmt.Errorf()
	fmt.Println("using fmt.Errorf()")

	res, err := divide(4, 0)
	if err != nil {
		fmt.Println("Error :", err)
	}
	fmt.Println(res)

	// using error interface
	fmt.Println("using error interface")

	sol, err := solution(5, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sol)

}