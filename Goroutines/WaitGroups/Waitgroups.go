package main

import (
	"fmt"
	"sync"
	"time"
)

func test(wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		fmt.Println("Hello World ", i)
		time.Sleep(time.Second * 2)
	}
	defer wg.Done()
}

func simplePrint(wg *sync.WaitGroup) {
	wg.Done()
	fmt.Println("Hello World")
	time.Sleep(time.Second * 1)
}

func main() {
	var wg sync.WaitGroup
	numberOfGoroutines := 5
	wg.Add(numberOfGoroutines)
	for i := 0; i < numberOfGoroutines; i++ {
		// Fast and Parallel Execution
		go simplePrint(&wg)
	}
	wg.Wait()
	time.Sleep(time.Second * 1)
	fmt.Print("done with execution")
}
