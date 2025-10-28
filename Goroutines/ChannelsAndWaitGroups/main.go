package main

import (
	"fmt"
	"sync"
	"time"
)

// here the channel is first filled then looped through

func makingEntriesIntoChannels(c *chan int, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		*c <- i
	}
	wg.Done()
	close(*c)
	fmt.Println("closing channel")
}

func main() {
	c := make(chan int, 5)
	var wg sync.WaitGroup
	wg.Add(1)
	go makingEntriesIntoChannels(&c, &wg)
	wg.Wait()
	// Scanning for Values to be entered into the channel
	for v := range c {
		fmt.Println("first Channel", v)
		time.Sleep(time.Second * 1)
	}
}
