package main

import "fmt"

func sumParis(nums []int) int {
	sum := 0

	for _, num := range nums {
		if num%2 == 0 {
			sum += num
		}
	}
	return sum
}

func main() {
	fmt.Println(sumParis([]int{1, 2, 3, 4, 5, 6}))
}
