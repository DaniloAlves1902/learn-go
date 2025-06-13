package main

import "fmt"

func invertSignal(number int) int {
	return number * -1
}

func invertSignalWithPointer(number *int) {
	*number = *number * -1
}

func main() {
	number := 3
	invertedNumber := invertSignal(number)
	fmt.Println(invertedNumber)

	newNumber := 12
	fmt.Println(newNumber)
	invertSignalWithPointer(&newNumber)
	fmt.Println(newNumber)
}
