package main

import (
	"fmt"
	"sync"
	"time"
)

func simplePrint(wg *sync.WaitGroup) {
	wg.Done()
	fmt.Println("Hello World")
	time.Sleep(time.Second * 5)
}

func main() {
	var wg sync.WaitGroup
	numberOfGoroutines := 5
	wg.Add(numberOfGoroutines)
	for i := 0; i < numberOfGoroutines; i++ {
		// Fast and Concurrent execution
		go simplePrint(&wg)
	}
	wg.Wait()
	time.Sleep(time.Second * 1)
	fmt.Print("done with execution")
}

// printing 5 Hello world strings will be completed within 1 second because we have 5 goroutines. main goroutine is waiting for 5 goroutines to be done, where each take milliseconds to be done.

// Output
// Hello World
// Hello World
// Hello World
// Hello World
// Hello World
// done with execution
