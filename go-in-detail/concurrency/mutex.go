package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter int
	var mutex sync.Mutex

	for range 1000 {
		go func() {
			mutex.Lock() // since we are trying to work with resource `counter`
			counter++
			mutex.Unlock()
		}()
	}
	mutex.Lock()
	fmt.Println("Expected :1000 , Executed:", counter)
	mutex.Unlock()
	// here we solve ONLY race condition
}
