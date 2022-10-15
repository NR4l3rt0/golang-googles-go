package main

import (
	"encoding/json"
	"exercises/Ninja11/person"
	"fmt"
	"log"
)

func main() {
	p1 := person.Person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last whishes?", "Never say never"},
	}

	bs, err := json.Marshal(p1)
	if err != nil {
		log.Fatalln(err)
		return
	}
	fmt.Println(string(bs))
}
