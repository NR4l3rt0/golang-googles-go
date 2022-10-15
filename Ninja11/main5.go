package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	fmt.Println(PositiveInt(num))
}

func PositiveInt(n int) bool {
	return n > 0
}
