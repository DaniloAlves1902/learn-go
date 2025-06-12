package main

import "fmt"

type user struct {
	name    string
	age     uint
	address address
}

type address struct {
	addressLine string
	number      uint
}

func main() {
	var user1 user
	user1.name = "Danilo"
	user1.age = 19

	user2 := user{"Sarah", 17, address{"rua 1", 10}}
	fmt.Println(user1)
	fmt.Println(user2)

	user3 := user{name: "João"}
	fmt.Println(user3)
}
