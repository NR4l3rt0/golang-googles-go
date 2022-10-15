package main

import (
	"fmt"
	"log"
)

type sqrtErr struct {
	lat  string
	long string
	err  error
}

func (se sqrtErr) Error() string {
	return fmt.Sprintf("Math error: %s, %s, %v", se.lat, se.long, se.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Exiting program")
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		//e := errors.New("Number cannot be less than 0")
		e := fmt.Errorf("Number cannot be less than 0. Value was: %f", f)
		return 0, sqrtErr{"1.234 N", "4.321 W", e}
	}
	return 42, nil
}
