package main

import "fmt"

func recoverExecution() {
	if r := recover(); r != nil {
		fmt.Println("execution successfully recovered")
	}
}

func division(numerator, denominator float32) bool {
	defer recoverExecution()

	if denominator == 0 {
		panic("cannot divide by 0")
	}

	division := numerator / denominator
	fmt.Printf("division: %.2f\n", division)
	return true
}

func main() {

	fmt.Println(division(7, 1))
	fmt.Println("post-execution")

}
