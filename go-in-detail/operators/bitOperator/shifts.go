package main

import "fmt"

func main() {
	// shift operators --> '<<' '>>'
	// num >> i
	// num = number
	// i = number of bits

	fmt.Println("shift operators --> '<<' '>>'")

	var a uint8 = 5 // in binary 5 = 101
	b := a >> 2
	// 101 >> 2 = 001
	// Right shift REMOVES 2 bits
	fmt.Println(b)
	// 001 = 1

    // formula : num / 2^1
    //           5 / 2^2 => 5/4 = 1

	var c uint8 = 8 // in binary 8 = 1000
	d := c << 4
	// 1000 << 4 = 10000000
	// LEFT shift ADDS 4 bits
	fmt.Println(d)
    // 10000000 = 128

    // formula : num x 2^i 
    //           8 x 2^4 => 8 x 16 = 128

}