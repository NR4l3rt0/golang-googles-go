package main

import (
	"fmt"
	"encoding/json"
	"os"
	"sort"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	Name string 
	Password string
	Nickname string
}

type ByNickname []user

func (a ByNickname) Len() int           { return len(a) }
func (a ByNickname) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByNickname) Less(i, j int) bool { return a[i].Nickname < a[j].Nickname }

func main() {
	fmt.Println("\n\nExercise 1\n")
	u1 := user{
		Name: `Peter`,
		Password: `1234`,
		Nickname: `ptr59`,
	}

	u2 := user{
		Name: `Claud`,
		Password: `4343`,
		Nickname: `c42`,
	}
	
	users := []user{u1,u2}
	bs, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
		return
	} 
	
	//fmt.Println(bs)
	fmt.Println(string(bs))

	fmt.Println("\n\nExercise 2\n")
	jsonUsers := `[{"Name":"Lucy","Password":"xxxyyy","Nickname":"luxy"},{"Name":"Mar","Password":"?hi","Nickname":"Sea09"}]`
	newUsers := []user{}
	json.Unmarshal([]byte(jsonUsers), &newUsers)
	if err != nil {
		fmt.Println(err)
		return
	} 
	fmt.Printf("%+v\n", newUsers)

	fmt.Println("\n\nExercise 3\n")
	//err = json.NewEncoder(os.Stdout).Encode(newUsers)
	err = json.NewEncoder(os.Stderr).Encode(newUsers)
	if err != nil {
		fmt.Println(err)
		return
	} 

	fmt.Println("\n\nExercise 4\n")
	xi := []int{3,24,4,41,1}
	xs := []string{`sandra`, `alberto`, `palmera`}
	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)
	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)

	fmt.Println("\n\nExercise 5\n")
	fmt.Println(users)
	sort.Sort(ByNickname(users))
	fmt.Println(users)

	fmt.Println("\n\nExercise 6 - Extra\n")
	for i, u := range users {
		xbs, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			fmt.Println(err)
			return
		} 
		users[i].Password = string(xbs)	
	}
	fmt.Println(users)
}
