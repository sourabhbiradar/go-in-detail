package main

import "fmt"

// func as type
type fType func(int) int // fType is of type func(int)int

func main() {
	// first-class citizens
	// higher-order func
	fmt.Println("\thigher-order func")

	// anonymous func
	fn := func(n int) {
		fmt.Println("\t'fn' anonymous func")
		fmt.Println("n :", n)
	}
	fn(5)

	// func as parameter
	fmt.Println("\tfunc paramFunc()")
	paramFunc(square, 5)

	// func as return value
	fmt.Println("\tfunc as return value")

	fmt.Println(returnFunc()(6)) // higher order func

	// func as type
	fmt.Println("\tfunc as type")
	fmt.Println(funcType()(11))

}

// func as parameter
func square(n int) (sq int) { // using this func as param
	sq = n * n
	return
}

func paramFunc(f func(int) int, num int) { // passing func square(n int) --> as param f func(int)
	fmt.Println("paramFunc()")
	fmt.Println("square of num :", f(num)) // f(num) --> square(n int)
}

// func as return value
func returnFunc() func(int) int {
	fmt.Println("returnFunc()")
	return square
}

// func as type
func funcType() fType { // fType --> func(int)int
	fmt.Println("funcType()")
	return square
}