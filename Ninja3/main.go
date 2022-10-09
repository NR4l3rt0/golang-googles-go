package main

import "fmt"

func main() {

	fmt.Println("\n\nExercise 1\n")
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("\n\nExercise 2\n")
	x := int('A')
	diff := int('Z') - x
	for i := 0; i <= diff; i++ {
		fmt.Println(x + i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U\n", x+i)
		}
	}

	fmt.Println("\n\nExercise 3\n")

	i := 1990
	for i <= 2022 {
		fmt.Println(i)
		i++
	}

	fmt.Println("\n\nExercise 4\n")

	j := 1990
	for {
		if j > 2022 {
			break
		}
		fmt.Println(j)
		j++
	}

	fmt.Println("\n\nExercise 5\n")

	for k := 10; k < 100; k++ {
		fmt.Println(k % 4)
	}

	fmt.Println("\n\nExercise 6\n")

	if a := false; !a {
		fmt.Println("That was a tricky question")
	}

	fmt.Println("\n\nExercise 7\n")

	if a := true; !a {
		fmt.Println("That was a tricky question")
	} else if 2 != 2 {
		fmt.Println("What was that?")
	} else {
		fmt.Println("That was another tricky question")
	}

	fmt.Println("\n\nExercise 8\n")

	switch {
	case false:
		fmt.Println("false")
	default:
		fmt.Println("default")
	}

	fmt.Println("\n\nExercise 9\n")

	favSport := "coding"
	switch favSport {
	case "basketball":
		fmt.Println("should not print")
	case "football":
		fmt.Println("should not print")
	case "running":
		fmt.Println("should not print")
	case "coding":
		fmt.Println("Coding should be a sport")
	default:
		fmt.Println("default")
	}

	fmt.Println("\n\nExercise 10\n")

	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
}
