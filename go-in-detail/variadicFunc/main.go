package main

import "fmt"

func variadicFunc(s ...string) { // paramName , 3 dots & type of param
	fmt.Println("func variadicFunc()")
	fmt.Println("s :", s)
}

func rangeParams(nums ...int) (sum int) {
	fmt.Println("func rangeParams()")
	for _, val := range nums {
		sum += val
	}
	return
}

func spreadOp(nums ...int) (sum int) {
	fmt.Println("func spreadOp()")
	for _, val := range nums {
		sum += val
	}
	return
}

func main() {
	// variadic func
	fmt.Println("\tvariadic func")

	variadicFunc("s", "l", "i", "c", "e") // multiple params

	// ranging over params
	fmt.Println("sum :", rangeParams(10, 20, 30))

	// using spread operator
	a := []int{-6, -4, -2}
	fn := spreadOp(a...)  
	fmt.Println("sum :", fn)

}