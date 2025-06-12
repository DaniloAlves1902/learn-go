package main

import "fmt"

func function1() {
	fmt.Println("executing function 1")
}

func function2() {
	fmt.Println("executing function 2")
}

func calculateAvarage(grade1 float32, grade2 float32) bool {

	defer fmt.Println("calculated average. The result will be returned")

	avarage := (grade1 + grade2) / 2
	if avarage >= 7 {
		return true
	}

	return false
}

func main() {
	//defer function2()
	//function1()

	fmt.Println(calculateAvarage(6, 8))
}
