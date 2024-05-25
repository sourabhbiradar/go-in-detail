package main

import (
	"fmt"
	"sync"
	"time"
)

func mmain() {
	var wg sync.WaitGroup
	wg.Add(5) // wait for 5 goR to finish

	// using for loop to create multiple goR at once
	for i := 0; i < 5; i++ {
		go worker(i, &wg) // run 5 goR
	}

	wg.Wait() // wait for all goR to finish

	fmt.Println("all goR finished")
}

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // decrement count

	fmt.Printf("worker %v started\n", id)

	time.Sleep(time.Second) // simulating some work

	fmt.Printf("worker %v finished\n", id)
}
