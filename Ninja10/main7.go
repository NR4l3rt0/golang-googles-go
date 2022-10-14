package main

import (
	"fmt"
	"sync"
	"runtime"
)

var wg sync.WaitGroup

func main() {

	c := make(chan int)
	myNum := 10
	for i := 0; i < 10; i++ {
		go func() {
			myNum--
			c <- myNum
		}()

	}

	for i := 0; i < 10; i++ {
		fmt.Printf("Num. goroutines: %d; channel says: %d\n", runtime.NumGoroutine(), <-c)
	}
	fmt.Println("About to exit")
}
