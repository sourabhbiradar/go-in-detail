package main

import "fmt"

func main() {
	// switch
	fmt.Println("\tswitch statement")

	a := 2 // try a := 5

	switch a {
	case 1:
		fmt.Println("first case")
	case 2:
		fmt.Println("second case")
	case 3:
		fmt.Println("third case") // since a == 3 'case 3' is executed
	default:
		fmt.Println("default case") // when 'a' niether matches any cases  'default' case is executed.
	}

	// switch 2
	fmt.Println("\tswitch 2")
	text := "India" // try text := "pyramid"

	switch x := len(text); x {
	case 4: // value
		fmt.Printf("%s is 4 letter word\n", text)
	case 5:
		fmt.Printf("%s is 5 letter word\n", text) // case 5 executes since length of text is 5
	default:
		fmt.Printf("%s is neither 4 nor 5 letter word\n", text)
	}

	// blank switch
	fmt.Println("\tblank switch")
	word := "India" // try word := "pyramid"

	switch x := len(text); { // no switching on value bt on condition
	case x == 4: // condition
		fmt.Printf("%s is 4 letter word\n", word)
	case x == 5:
		fmt.Printf("%s is 5 letter word\n", word) // case 5 executes since length of text is 5
	default:
		fmt.Printf("%s is neither 4 nor 5 letter word\n", word)
	}

	// multiple values/conditions
	fmt.Println("\tmultiple values/conditions")

	switch num := 2; num {
	case 1, 2: // multiple values
		fmt.Println("num is either 1 or 2")
	case 3, 4:
		fmt.Println("num is either 3 or 4")
	default:
		fmt.Println("num is neither 1, 2,3 nor 4")
	}

	// fallthrough
	fmt.Println("\tfallthrough")
	switch num := 1; num { // try num := 2
	case 1:
		fmt.Println("num is 1")
		fallthrough // case 2 will also be executed
	case 2:
		fmt.Println("num is 2")
		fallthrough // case 3 will also be executed
	case 3:
		fmt.Println("num is 3")
	case 4: // not executed as case 3 does NOT have 'fallthrough'
		fmt.Println("num is 4")
	default:
		fmt.Println("num is neither 1, 2,3 nor 4")
	}

}