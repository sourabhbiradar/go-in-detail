package main

import "fmt"

func main() {
	// control flow statements

	// break
	fmt.Println("\tbreak")

	for true { // infinite loop
		fmt.Println("just executes once and exits ")
		break // without break loop would iterate forever
	}

	// break/exit condition
	for i := 1; i < 5; i++ {
		fmt.Println("i :", i)
		if i == 2 { // code exits when 'i' value reaches '3'
			break // without break loop would iterate till i < 5
		}
	}
	fmt.Println("Din't execute i = 3 and i = 4")

	// contiinue
	fmt.Println("\tcontinue")

	for i := 1; i <= 5; i++ {

		if i == 2 {
			fmt.Println("skippin iteration")
			continue // without continue even i : 2 would be executed
		}
		fmt.Println("i :", i) // when i = 2 its not executed
	}

	// goto
	fmt.Println("\tgoto")
	// entry: fmt.Println("at entry label")

	a := 0  
    // if 'a' is anything but '0' both "a :" and "ending program" will be executed
    // try a := 1
	if a == 0 {
		goto end
	}
	fmt.Println("a :", a) // will NOT be executed
end: // end label
	fmt.Println("ending program")

	// label
	fmt.Println("\tlabel")

outerLoop: // label outerLoop
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				break outerLoop // skipped every iteration after i = 1 j = 1
			}
			fmt.Println(i, j)
		}
	}

}