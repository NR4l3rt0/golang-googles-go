package main

import "fmt"

type person struct {
	first      string
	last       string
	favFlavors []string
}

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	fmt.Println("\n\nExercise 1\n")
	p1 := person{
		first:      `Peter`,
		last:       `Watson`,
		favFlavors: []string{`cherry`, `banana`, `lemon`},
	}
	p2 := person{
		first:      `Paul`,
		last:       `Bird`,
		favFlavors: []string{`melon`, `apple`},
	}
	fmt.Printf("%T\t %+v\n", p1, p1)
	fmt.Printf("%T\t %+v\n", p2, p2)

	for _, v := range p1.favFlavors {
		fmt.Println(v)
	}
	for _, v := range p2.favFlavors {
		fmt.Println(v)
	}


	fmt.Println("\n\nExercise 2\n")
	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}
	for k, v := range m {
		fmt.Println(k, ":", v)
	}


	fmt.Println("\n\nExercise 3\n")
	t := truck{
		vehicle: vehicle{
			doors: 2,
			color: "blue",
		},
		fourWheel: false,
	}
	s := sedan{
		vehicle: vehicle{
			doors: 2,
			color: "black",
		},
		luxury: true,
	}

	fmt.Println(t)
	fmt.Println(t.color, t.fourWheel)
	fmt.Println(s)
	fmt.Println(s.color, s.luxury)


	fmt.Println("\n\nExercise 4\n")
	x := struct {
		id  int
		src string
		dst string
	}{
		id:  1234,
		src: "Malaga",
		dst: "New York",
	}

	fmt.Printf("%T\t %+v\n", x, x)

	y := struct {
		sedan
		connection map[string][]string
	}{
		sedan: sedan{
			luxury: true,
		},
		connection: map[string][]string{
			"trip1": []string{"Malaga", "New York"},
			"trip2": []string{"Berlin", "Paris"},
		},
	}

	fmt.Printf("%T\t %+v\n", y, y)
}
