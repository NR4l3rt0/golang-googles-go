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

	bs, err := toJSON(p1)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println(string(bs))
}

func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		myErr := fmt.Errorf("Error found while marshal operation: %v", err)
		return []byte{}, myErr
	}
	return bs, nil
}
