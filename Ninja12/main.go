package main

import (
	"exercise/Ninja12/dog"
	"fmt"
)

func main() {

	var age int
	fmt.Print("Please, insert an age: ")
	fmt.Scan(&age)

	d := dog.Dog{
		Name: `Tom`,
		Age:  dog.Years(age),
	}
	fmt.Printf("\nYour dog %s is %d years old\n", d.Name, d.Age)
}
