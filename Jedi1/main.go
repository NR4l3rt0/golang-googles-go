package main

import "fmt"

// Package scope
var a int
var b string
var c bool

var d = 42
var e = "James Bond"
var f = true

func main() {
	fmt.Printf("Exercise 1")
	// Local scope
	x, y, z := 42, "James Bond", true
	fmt.Printf("x => %v, y => %v, z => %v\n", x, y, z)
	fmt.Printf("x => %d\n", x)
	fmt.Printf("y => %s\n", y)
	fmt.Printf("z => %t\n", z)

	fmt.Printf("\n\nExercise 2\n")
	fmt.Printf("Value => %d; Type: %T\n", a, a)
	fmt.Printf("Value => %s; Type: %T\n", b, b)
	fmt.Printf("Value => %t; Type: %T\n", c, c)

	fmt.Printf("\n\nExercise 3\n")
	s := fmt.Sprintf("%d %q %t", d, e, f)
	fmt.Printf("%v\n", s)
}
