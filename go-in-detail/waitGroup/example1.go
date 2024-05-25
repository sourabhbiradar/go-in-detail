package main

import (
	"fmt"
	"sync"
)

func helloPrinter(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("hello")
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	go helloPrinter(&wg)
	go helloPrinter(&wg)

	wg.Wait()
}
