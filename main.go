package main

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
	u := NewUser()
	// fmt.Println("Talkが呼ばれる前" + u.name)
	Talk(u)
	// fmt.Println("Talkが呼ばれたあと" + u.name)
}
