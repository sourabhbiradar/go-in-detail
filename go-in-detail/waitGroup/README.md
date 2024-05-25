# waitGroup

`WaitGroup` struct exported from `sync` package.

`sync.WaitGroup` is used to wait for collections of go routines to finish executing.

This is useful for coordinating the completion of concurrent tasks.

Prevents go routines leak. 

WAitGroup is `memory safe counter`;
Add() : increments count of go routines to wait for;
Done() : decrements count once goR is executed.

steps :

var wg sync.WaitGroup   // initialize;
wg.Add(n)  // n = number of go routines;
wg.Done()  // indicates a go routines is finished executing;
wg.Wait()  // waits until all `n` is zero;
  