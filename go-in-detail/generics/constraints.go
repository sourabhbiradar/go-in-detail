package main

import (
	"fmt"
)

func main() {
	Any([]string{"a", "b"}, "c", "")
	// Any([]string{"a", "b"}, 8, 3.14)
	// since T is set to string following values have to be of string type

	Comp([]int{1, 2, 3}, 3)

	fmt.Println(ID.Equal(3, 3))

	tiltStr([]str{"a", "b"}) // [T ~string] type str string

	union(1, 2)
	union("a", "b") // a,b can only be of type int or string

	genFn(i(0))
	genFn(f(0))

}

// [any]
func Any[T any](list []T, num T, float T) {
	fmt.Printf("%v %v %v", list, num, float)
}

// [comparable]
func Comp[T comparable](list []T, val T) bool {
	for _, v := range list {
		if v == val {
			return true
		}
	}
	return false
}

// [interface]
type ID int

func (id ID) Equal(other ID) bool {
	return id == other
}

type Intfc[T any] interface {
	Equal(other T) bool
}

// [~type]
type str string

func tiltStr[T ~string](s []T) {
	fmt.Println(s)
}

// [T int|uint]
func union[T int | string](a, b T) {

	fmt.Println(a, b)
}

// combining constraints

type Group interface {
	~int | ~float32
	fn1(a i)
	fn2(b f)
}

type i int
type f float32

// generic function
func genFn[T Group](t T) {
	t.fn1(4)
	t.fn2(3.4444444)
}
func (i) fn1(a i) {
	fmt.Println(a)
}
func (f) fn2(b f) {
	fmt.Printf("%.*f\n", 3, b)
}
func (f) fn1(b i) {
	fmt.Println(b)
}
func (i) fn2(b f) {
	fmt.Printf("%.*f\n", 3, b)
}
