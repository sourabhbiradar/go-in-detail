# concurrency

Dividing & executing program is called concurrency.

Concurrency may seem like parallelism but
Concurrency != parallelism

Increases performance of program by using computer resources efficiently.

`Program >> processor >> thread >> Go Routines`

Problem with concurrency & Go routines:
1. Can't determine the outcome
2. Some go routines may not be even executed
3. Race condition where multiple go routines compete for same resource.

To solve all this we use time.Sleep()

But for heavy computations we use `ATOMICITY` where there are lot of go routines

We use `sync.Mutex` to lock & unlock resources so that only one go routine is accessing resource.

More abt sync.Mutex in later codes.
