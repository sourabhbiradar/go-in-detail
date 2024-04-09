package main

import "fmt"

func main() {
	// string in depth

	str := "Pirates"

	fmt.Println("index 4 :", str[4])             // 116 --> ASCII code for 't'
	conv := string(str[4])                       // type casting 116 to string
	fmt.Println("converted to string -->", conv) // t

	fmt.Println("index 2 to 4 :", string(str[2:5])) // rat

	// using []rune() -->> RECOMMANDED
	kan := "ಸೌರಭ್"
	fmt.Println("Tryin to access 'ರ' :", string(kan[1])) // can NOT

	r := []rune(kan)      // or r := []byte(kan)
	fmt.Println("ASCII code for every elem in kan :", r)
	fmt.Println(string(r[2])) // ರ

	// so when dealin with unicodes or other lanuages its BETTER to convert string into rune/byte than type cast them.

}
