package main

import "fmt"

const (
	num     int = 3       // typed const
	unTyped     = "Peter" // untyped const
)

const (
	year1 = 2022
	year2 = (year1 - iota)
	year3 = (year1 - iota)
	year4 = (year1 - iota)
)

func main() {

	fmt.Printf("Exercise 1\n")
	x := 42
	fmt.Printf("d => %d\n", x)
	fmt.Printf("b => %b\n", x)
	fmt.Printf("x => %#x\n", x)

	fmt.Printf("\n\nExercise 2\n")
	x = 42
	fmt.Println(x == 42)
	fmt.Println(x != 42)
	fmt.Println(x > x/3)

	fmt.Printf("\n\nExercise 3\n")
	fmt.Println(num, unTyped)
	fmt.Printf("%T\t%T\t", num, unTyped)

	fmt.Printf("\n\nExercise 4\n")
	y := 2
	fmt.Printf("%d\t %b\t %#X\n", y, y, y)
	z := y << 1
	fmt.Printf("%d\t %b\t %#X\n", z, z, z)

	fmt.Printf("\n\nExercise 5\n")
	rs := `This is
		a raw String 
		in Go`
	fmt.Println(rs)

	fmt.Printf("\n\nExercise 6\n")
	fmt.Println(year1, year2, year3, year4)
}
