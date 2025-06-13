package main

import "fmt"

type user struct {
	name string
	age  int8
}

func (u user) save() {
	fmt.Printf("user %s saved successfully\n", u.name)
}

func (u user) ofLegalAge() bool {
	return u.age >= 18
}

func (u *user) haveABirthday() {
	u.age++
}

func main() {

	user1 := user{"Danilo", 19}
	user1.save()

	fmt.Println("is of legal age:", user1.ofLegalAge())
	fmt.Println(user1.age)
	
	user1.haveABirthday()
	fmt.Println(user1.age)

}
