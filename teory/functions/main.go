package main

import "fmt"

func main() {
	sum := sum(1, 1)
	fmt.Println(sum)

	var subtract = func(n1 int16, n2 int16) int16 {
		return n1 - n2
	}

	sub := subtract(1, 1)
	fmt.Println(sub)

	mult := mult(2, 2)
	fmt.Println(mult)

	var div = func(n1, n2 int16) int16 {
		return n1 / n2
	}
	fmt.Println(div(10, 2))

	valueSum, _, _, _ := calculator(10, 20)
	fmt.Println("sum is", valueSum)

	_, valueSub, _, _ := calculator(10, 5)
	fmt.Println("subtract is", valueSub)

	_, _, valueMult, _ := calculator(2, 2)
	fmt.Println("multiplicate id", valueMult)

	_, _, _, valueDiv := calculator(10, 2)
	fmt.Println("division is", valueDiv)

}

func sum(n1 int16, n2 int16) int16 {
	return n1 + n2
}

func mult(n1, n2 int16) int16 {
	return n1 * n2
}

func calculator(n1, n2 float32) (float32, float32, float32, float32) {
	sum := n1 + n2
	sub := n1 - n2
	mult := n1 * n2
	div := n1 / n2
	
	return sum, sub, mult, div
}
