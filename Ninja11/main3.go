package main

import (
	"errors"
	"fmt"
	"log"
)

type customErr struct {
	Name   string
	Number int
	Err    error
}

func (ce customErr) Error() string {
	return fmt.Sprintf("Error name: %s; exit code: %d; msg: %s", ce.Name, ce.Number, ce.Err)
}

func main() {

	err := customErr{
		Name:   "custom",
		Number: 5,
		Err:    errors.New("Situation not handled"),
	}

	foo(err)
	fmt.Println("Exiting program")
}

func foo(e error) {
	log.Printf("Type %T: ", e)
	log.Println(e)

	// Assertion (if not, cannot access underline fields)
	log.Println(e.(customErr).Err, "!!!")
	//log.Println(e.Err, "!!!") <- error type doesn't have that field
}
