package main

import (
	"fmt"
	"sync"
	"time"
)

// here the channel is first filled then looped through

func makingEntriesIntoChannels(c *chan int, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		fmt.Println("inserting into channel", i)
		*c <- i
	}
	wg.Done()
	close(*c)
	fmt.Println("closing channel")
}

func main() {

	// making Channel with some space in it
	c := make(chan int, 5)
	var wg sync.WaitGroup
	wg.Add(1)
	go makingEntriesIntoChannels(&c, &wg)
	wg.Wait()
	// Scanning for Values to be entered into the channel
	for v := range c {
		fmt.Println("Reading Channel", v)
		time.Sleep(time.Second * 1)
	}
}

// Output
// inserting into channel 0
// inserting into channel 1
// inserting into channel 2
// inserting into channel 3
// inserting into channel 4
// closing channel
// Reading Channel 0
// Reading Channel 1
// Reading Channel 2
// Reading Channel 3
// Reading Channel 4
