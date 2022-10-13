package main

import (
	"fmt"
)


func main(){
	 c := gen()
	receive(c)
	fmt.Printf("About to exit\n")
}


func gen() <-chan int {
	c := make(chan int)

	go func(){
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}


func receive(cr <-chan int) {
	for v := range cr {
		fmt.Println(v)
	}
}
