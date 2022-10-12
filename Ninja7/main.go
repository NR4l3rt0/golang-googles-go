package main

import "fmt"

type person struct{
	name string	
}

func changeMe(p *person, newName string){
	//(*p).name = newName
	p.name = newName
} 

func main() {
	fmt.Println("\n\nExercise 1\n")
	x := 42
	fmt.Printf("%T %d\n", x, x)
	fmt.Printf("%T %v\n", &x, &x)

	fmt.Println("\n\nExercise 2\n")
	p := person{`Pet`}
	fmt.Printf("%T %+v\n", p, p)
	fmt.Printf("%T %v\n", &p, &p)
	changeMe(&p, `Chris`)
	fmt.Printf("%T %+v\n", p, p)
	fmt.Printf("%T %v\n", &p, &p)
	changeMe(&p, `Tom`)
	fmt.Printf("%T %+v\n", p, p)
	fmt.Printf("%T %v\n", &p, &p)

	fmt.Println("\n\nExercise 3 - Extra\n")
	p2 := &person{`Pet`}
	fmt.Printf("%T %+v\n", p2, p2)
	fmt.Printf("%T %v\n", *p2, *p2)
	changeMe(p2, `Chris`)
	fmt.Printf("%T %+v\n", p2, p2)
	fmt.Printf("%T %v\n", *p2, *p2)
	changeMe(p2, `Tom`)
	fmt.Printf("%T %+v\n", p2, p2)
	fmt.Printf("%T %v\n", *p2, *p2)
}
