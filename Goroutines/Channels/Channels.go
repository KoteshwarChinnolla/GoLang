package main

import (
	"fmt"
	"time"
)

func makingEntriesIntoChannels(c *chan int) {
	for i := 0; i < 5; i++ {
		*c <- i
		*c <- i * 2
	}
	fmt.Println("closing channel")
	close(*c)
}

func main() {
	c := make(chan int, 2)
	go makingEntriesIntoChannels(&c)
	// Scanning for Values to be entered into the channel
	for v := range c {
		fmt.Println("first Channel", v)
		time.Sleep(time.Second * 1)
	}
	// The channel is closed here so it will not print anything
	for v := range c {
		fmt.Println("Second Channel", v)
		time.Sleep(time.Second * 1)
	}
}
