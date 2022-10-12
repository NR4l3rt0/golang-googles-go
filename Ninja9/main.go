package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup
var wg2 sync.WaitGroup
var mu sync.Mutex

type person struct {
	name string
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Print(p.name)
}

func saySomething(h human) {
	h.speak()
	fmt.Println(` - Only pointers can be humans!`)
}
func main() {
	fmt.Println("\n\nExercise 1\n")
	fmt.Println(`Num. CPU: `, runtime.NumCPU())
	fmt.Println(`Active Go routines: `, runtime.NumGoroutine())

	wg.Add(1)
	go func() {
		fmt.Println("Go rutine 1")
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		runtime.Gosched()
		fmt.Println("Go rutine 2")
		wg.Done()
	}()

	fmt.Println(`Active Go routines: `, runtime.NumGoroutine())

	wg.Wait()
	fmt.Println(`Active Go routines: `, runtime.NumGoroutine())
	fmt.Println("\nExercise 1 exited!")

	fmt.Println("\n\nExercise 2\n")

	p1 := &person{`Joe`}
	p2 := person{`Tini`} //<- not working because it's not a pointer
	saySomething(p1)
	//saySomething(p2) - Cheating a bit for it to work
	saySomething(&p2)

	fmt.Println("\n\nExercise 3 and 4\n")

	counter := 0
	wg2.Add(200)
	for i := 0; i < 200; i++ {
		go func() {
			mu.Lock()
			fmt.Println(counter)
			v := counter
			//runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			wg2.Done()
		}()
	}

	wg2.Wait()
	fmt.Println("Final counter: ", counter)

	fmt.Println(`Exiting exercise 5`)
	var atomCounter int64 
	wg2.Add(200)
	for i := 0; i < 200; i++ {
		go func() {
			
			fmt.Println(atomic.LoadInt64(&atomCounter))
			atomic.AddInt64(&atomCounter, 1)
			wg2.Done()
		}()
	}

	wg2.Wait()
	fmt.Println("Final atomic counter: ", atomic.LoadInt64(&atomCounter))

	fmt.Println(`Exercise 6`)
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}
