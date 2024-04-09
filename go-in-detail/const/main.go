package main

import "fmt"

const a int32 = 10 // global const

func main() {
	// const --> constants

	const x = 8 // its ok to declared and not use 'x'

	var i int32 = a     // can be assigned to 'i' as both are of type "int32"
    fmt.Println("i :", i)

   var j int = a       // can't assign since 'j' is of type int64 and 'a' is "int32"
   fmt.Println("j :", j)
    // fix this to unlock code ;)
    //hint : match types or type cast 'a'

	fmt.Println("'a' is global const :", a)

    const n = "string"    // untyped const
    const m int = 4       // typed const

}