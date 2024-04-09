package main

import "fmt"

func main() {
	// anonymous structs

	player := struct {
		name string
		runs int
		wkts int
	}{
		name: "Abc",
		runs: 56,
		wkts: 2,
	}

	fmt.Printf("player stats : %+v\n", player)
	fmt.Println("player name :", player.name)

	// we can NOT use anonymous struct for various instences bt standard struct can be used.

}
