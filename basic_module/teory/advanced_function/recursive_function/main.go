package main

import "fmt"

func fibonacci(n uint) uint {
	if n <= 1 {
		return n
	}

	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {

	n := uint(8)

	for i := uint(0); i < n; i++ {
		fmt.Print(fibonacci(i), ", ")
	}

	//fmt.Println(fibonacci(n))

}
