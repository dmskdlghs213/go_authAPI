package main

import "fmt"

type User struct {
	name string
	age  string
}

func NewUser() *User {
	u := &User{
		name: "masa",
		age:  "21",
	}
	return u
}

func Talk(u *User) {
	u.name = "take"
}

func main() {
	fmt.Println("新しいユーザーが喋ります")
	u := NewUser()
	fmt.Println("Talkが呼ばれる前" + u.name)
	Talk(u)
	fmt.Println("Talkが呼ばれたあと" + u.name)
}
