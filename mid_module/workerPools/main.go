package main

import "fmt"

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	return fibonacci(n-1) + fibonacci(n-2)
}

func worker(tasks <-chan int, result chan<- int) {
	for n := range tasks {
		result <- fibonacci(n)
	}
}

func main() {

	tasks := make(chan int, 45)
	results := make(chan int, 45)

	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)

	for i := 0; i < 45; i++ {
		tasks <- i
	}

	close(tasks)

	for j := 0; j < 45; j++ {
		result := <-results
		fmt.Println(result)
	}

}
