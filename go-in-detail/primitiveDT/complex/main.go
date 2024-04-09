package main

import "fmt"

func main() {
	// complex data type

	var sComplex complex64 = complex(3.123456, 6.0987654)
	fmt.Println(sComplex) // (3.123456+6.0987654i)

	var bigCmplx complex128 = complex(5.0123456789101112, 2.0987654321)
	fmt.Println(bigCmplx) // (5.012345678910111+2.0987654321i)

	var cmplx complex64
	cmplx = complex(2, 3)
	fmt.Println(cmplx) // (2+3i)

	fmt.Println(real(cmplx))  // 2 --> real part
	fmt.Println(imag(cmplx))  // 3 --> imainar part

}