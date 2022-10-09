package main

import "fmt"

func main() {

	fmt.Printf("\nExercise 1\n\n")
	/*
	   var arr [5]int
	   	arr[0] = 22
	   	arr[1] = 3
	   	arr[2] = 42
	   	arr[3] = 5
	   	arr[4] = 4255
	*/
	arr := [5]int{22, 3, 42, 5, 4255}
	fmt.Printf("Type %T, Value: %v\n", arr, arr)
	for i, v := range arr {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}

	fmt.Printf("\nExercise 2\n\n")
	xi := []int{22, 3, 42, 5, 4255, 1, 2, 3, 4, 5}
	fmt.Printf("Type %T, Value: %v\n", xi, xi)
	for i, v := range xi {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}

	fmt.Printf("\nExercise 3\n\n")
	xx1 := xi[:4]
	xx2 := xi[5:9]
	fmt.Printf("Type %T, Value: %#v\n", xx1, xx1)
	fmt.Printf("Type %T, Value: %#v\n", xx2, xx2)

	fmt.Printf("\nExercise 4\n\n")
	slc := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Printf("Type %T, Value: %#v\n", slc, slc)

	slc = append(slc, 52)
	fmt.Printf("Type %T, Value: %#v\n", slc, slc)

	slc = append(slc, 53, 54, 55)
	fmt.Printf("Type %T, Value: %#v\n", slc, slc)

	y := []int{56, 57, 58, 59, 60}
	slc = append(slc, y...)
	fmt.Printf("Type %T, Value: %#v\n", slc, slc)

	fmt.Printf("\nExercise 5\n\n")
	slcDel := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Printf("Type %T, Value: %#v\n", slcDel, slcDel)

	slcDel = append(slcDel[:3], slcDel[6:]...)
	fmt.Printf("Type %T, Value: %#v\n", slcDel, slcDel)

	fmt.Printf("\nExercise 6\n\n")
	fixSlice := make([]int, 0, 50)
	fmt.Printf("Type %T, len: %d, cap: %d\n", fixSlice, len(fixSlice), cap(fixSlice))

	for i := 0; i < 50; i++ {
		fixSlice = append(fixSlice, i)
	}
	fmt.Printf("Type %T, len: %d, cap: %d\n", fixSlice, len(fixSlice), cap(fixSlice))
	fmt.Println(fixSlice)

	fmt.Printf("\nExercise 7\n\n")
	x1 := []string{"James", "Bond", "Shaken, not stirred"}
	x2 := []string{"Miss", "Moneypenny", "Hello, James"}
	xSlice := [][]string{x1, x2}

	for _, v := range xSlice {
		for _, innerVal := range v {

			fmt.Println(v, ":", innerVal)
		}
	}

	fmt.Printf("\nExercise 8 and 9\n\n")
	//m := map[string][]string{ `k` : []string{`vals`,...}}
	m := make(map[string][]string)
	m["bond_james"] = []string{"Shaken, not stirred", `Martinis`, `Women`}
	m["moneypenny_miss"] = []string{"James Bond", `Being evil`, `Sunset`}

	for k, v := range m {
		fmt.Println(k, ":", v)
		for _, v2 := range v {
			fmt.Println("\t-", v2)
		}
	}

	fmt.Printf("\nExercise 10\n\n")
	if _, ok := m["bond_james"]; ok {
		delete(m, "bond_james")
	}

	for k, v := range m {
		fmt.Println(k, ":", v)
		for _, v2 := range v {
			fmt.Println("\t-", v2)
		}
	}
}
