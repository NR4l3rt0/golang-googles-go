package main

import (
	"fmt"
)

func main(){
	c := make(chan int)
	go func(){
		send(c, 42)
	}()
	fmt.Print(receive(c))
	fmt.Printf("---------\n")
	fmt.Printf("c:\t%T\n", c)
}


func send(cs chan<- int, n int) {
	cs <- n
}

func receive(cr <-chan int) string{
	return fmt.Sprintln(<- cr)
}
