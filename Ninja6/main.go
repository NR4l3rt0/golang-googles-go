package main

import (
	"fmt"
	"math"
)

func foo() int {
	return 42
}

func bar() (int, string) {
	return foo(), "Hello"
}

func foo2(ints ...int) int {
	sum := 0
	for _, v := range ints {
		sum += v
	}
	return sum
}

func bar2(ints []int) int {
	sum := 0
	return func(xi []int) int {
		for _, v := range xi {
			sum += v
		}
		return sum
	}(ints)
}

func myDefer() {
	fmt.Println(`I'm being deferred`)
}

func myDefer2() {
	fmt.Println(`I'm being deferred; but I go out first`)
}

func normalFlow() {
	fmt.Println(`Just in time`)
}

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() string {
	return fmt.Sprintf(`I'm %s %s and I'm %d years old`, p.first, p.last, p.age)
}

type square struct {
	side int
}

type circle struct {
	radius float64
}

func (s square) calcArea() float64 {
	return float64(s.side * s.side)
}

func (c circle) calcArea() float64 {
	return math.Pi * c.radius * c.radius
}

type shape interface {
	calcArea() float64
}

func info(calc []shape) {
	for _, v := range calc {
		switch v.(type) {
		case square:
			fmt.Println(`Square says: `, v.calcArea())
		case circle:
			fmt.Println(`Circle says: `, v.calcArea())
		default:
			fmt.Println(`I'm confused`)
		}
	}
}

func isPeterHere(s string) bool {
	return s == "Peter"
}

func lookingFor(f func(s string) bool, xs []string) string {
	for _, v := range xs {
		if f(v) {
			return `Is there!`
		}
	}
	return `There is no Peter`
}

func myClosureGreeter(s string) func() string {
	i := 0
	fullName := "Bernard " + s

	return func() string {
		i++
		return fmt.Sprintln("I say hello to", fullName, i, "time(s)")
	}
}

func main() {
	fmt.Println("\n\nExercise 1\n")
	fmt.Println(foo())
	fmt.Println(bar())

	fmt.Println("\n\nExercise 2\n")
	myInts := []int{1, 2, 3, 4}
	sum1 := foo2(myInts...)
	fmt.Println(sum1)

	sum2 := bar2(myInts)
	fmt.Println(sum2)

	fmt.Println("\n\nExercise 3\n")
	defer myDefer()
	defer myDefer2()
	normalFlow()

	fmt.Println("\n\nExercise 4\n")
	p := person{
		first: `John`,
		last:  `Wayne`,
		age:   93,
	}
	fmt.Println(p.speak())

	fmt.Println("\n\nExercise 5\n")
	c := []shape{
		square{side: 4},
		circle{radius: 2},
	}

	info(c)
	fmt.Println()

	fmt.Println("\n\nExercise 6\n")
	// Anonymous func already used

	fmt.Println("\n\nExercise 7\n")
	f := func(i int, s string) string {
		return fmt.Sprintf(`%d more days, little %s`, i, s)
	}

	fmt.Println(f(42, `friend`))
	fmt.Println()

	fmt.Println("\n\nExercise 8\n")
	f2 := func(i int, s string) func() string {
		return func() string {
			return fmt.Sprintf(`%d more days, little %s`, i, s)
		}
	}
	fmt.Printf("%T\n", f2)
	fmt.Println(f2(33, `cat`)())

	fmt.Println("\n\nExercise 9\n")
	xs := []string{`Tobias`, `Lana`, `Peter`}
	xs2 := []string{`Tobias`, `Lana`}
	fmt.Println("Expected true, result: ", lookingFor(isPeterHere, xs))
	fmt.Println("Expected false, result: ", lookingFor(isPeterHere, xs2))

	fmt.Println("\n\nExercise 10\n")
	clo1 := myClosureGreeter("Shaw")
	clo2 := myClosureGreeter("Punch")
	fmt.Println(clo1())
	fmt.Println(clo1())
	fmt.Println(clo1())
	fmt.Println(clo1())
	fmt.Println(clo2())
	fmt.Println(clo2())
	fmt.Println(clo1())
	fmt.Println(clo2())
	fmt.Println()
}
