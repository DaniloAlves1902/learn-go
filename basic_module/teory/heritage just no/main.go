package main

import "fmt"

type person struct {
	name   string
	age    uint
	gender string
}

type student struct {
	person
	course     string
	university string
}

func main() {
	p := person{"Danilo", 19, "masculine"}
	fmt.Println(p)

	s := student{p , "ADS", "GRAN"}
	fmt.Println(s)

}
