package main

import (
	"fmt"
)

func main(){
	//Solution 1 with buffers | c := make(chan int, 1)
	c := make(chan int)
	go func(){
		c <- 42
	}()
	fmt.Println(<-c)
	fmt.Println("Exiting")
}
