package main

import "fmt"

func main() {

	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i + 1
		}
		close(c)
	}()

	/*
		for {
			if v, ok := <-c; ok {
				fmt.Println(v)
				continue
			}
			return
		}
	*/
	for v := range c {
		fmt.Println(v)
	}
	fmt.Println("About to exit")
}
