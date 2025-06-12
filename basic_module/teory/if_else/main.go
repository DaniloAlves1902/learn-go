package main

import "fmt"

func main() {
	number := -10

	if number > 15 {
		fmt.Println("is greater than 15")
	} else {
		fmt.Println("is less than or equal to 15")
	}

	// if init
	if otherNumber := number; otherNumber > 0 {
		fmt.Println("is greater than 0")
	} else if number <= 10{
		fmt.Println("is less than or equal to 10")
	} else {
		fmt.Println("between 0 and 10")
	}

}
