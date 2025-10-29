package main

import (
	"fmt"
	"time"
)

func printNumber(i int) {
	fmt.Println(i)
	time.Sleep(time.Second * 1)
}

func main() {
	limiters := make(chan int, 5)
	for i := 0; i < 5; i++ {
		limiters <- i
		go func() {
			printNumber(i)
			<-limiters
		}()
	}
}
