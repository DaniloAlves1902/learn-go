package main

import "fmt"

func main() {
	ch := make(chan string, 2)

	ch <- "hello, go!"
	ch <- "programing in GO"

	msg := <-ch
	msg2 := <-ch
	fmt.Println(msg)
	fmt.Println(msg2)
}
