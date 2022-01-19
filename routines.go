package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	// Go Routines are scheduled by Go Runtime
	// Go Routines are lighter wight than OS threads
	// Fewer context switching
	// Go parallelism is handled by go runtime
	// Go layers Go Routines on top of threads, when a GoRoutine blocks (i.e Working on a network IO) Go swaps it out
	// for another Go Routine but running on the same thread.
	// Even with Go Routines, there's gonna be block conditions needing a new thread, its just that happens way less often
	// Go concurrency model uses the Actor model. Go routines communicate and share data using channels

	
	runtime.GOMAXPROCS(2)

	var waitGrp sync.WaitGroup
	waitGrp.Add(2)

	//Adding () at the end of the function it becomes a self-executed funcion
	go func() {
		defer waitGrp.Done() // Reports to the WaitGrp when the execution of the current routine it's done

		time.Sleep(5 * time.Second)
		fmt.Println("Hello")
	}()

	go func() {
		defer waitGrp.Done() // Reports to the WaitGrp when the execution of the current routine it's done

		fmt.Println("World")
	}()

	waitGrp.Wait()

	//Channels
	//Unbuffered channels can't hold data, any go routine putting data 1 to one blocks until there's a receiver on the other end forcing a sync behavior
	//Buffered channels can hold data, they can put data on chanels and don't care if there was any receiver on the other end -> async behavior

	//creating unbuffered channels
	myUnBufferedChann := make(chan int)

	//creating buffered channels
	myBufferedChann := make(chan int, 5) // We are creating a channel that can hold up to 5 integers

}
