package main

import "fmt"

func simpleCalculator(number1, number2 int) (sum int, sub int) {
	sum = number1 + number2
	sub = number1 - number2
	return
}

func main() {

	sum, sub := simpleCalculator(10, 20)

	fmt.Println(sum, sub)

}
