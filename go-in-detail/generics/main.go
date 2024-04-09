// package main

// import "fmt"

// func main() {
// 	// Generics
// 	fmt.Println("\tGenerics")

// 	fmt.Printf("Has a %v\n", Has([]string{"a", "b"}, "a"))
// 	fmt.Printf("Has c %v\n", Has([]string{"a", "b"}, "c"))
// 	fmt.Printf("Has 2 %v\n", Has([]int{1, 2, 3}, 2))

// 	fmt.Printf("MyList %v\n", NewEmptyList[int]())

// 	MultiTypeParams(1, 2, "b", 3)
// }

// // func Has(list []string, value string) bool {
// // 	for _, v := range list {
// // 		if v == value {
// // 			return true
// // 		}
// // 	}
// // 	return false
// // }

// func Has[T comparable](list []T, value T) bool {
// 	for _, v := range list {
// 		if v == value {
// 			return true
// 		}
// 	}
// 	return false
// }

// func NewEmptyList[T any]() []T {
// 	return make([]T, 0)
// }

// // multiple type parameters

// func MultiTypeParams[A, B any, C ~int](a1, a2 A, b B, c C) {
// 	fmt.Printf("%v %v %v %v", a1, a2, b, c)
// }


