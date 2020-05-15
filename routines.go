package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {

	//go parallelism is handled by go runtime
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
}
