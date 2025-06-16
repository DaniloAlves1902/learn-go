package main

import (
	"fmt"
	"time"
)

// goroutines
func main() {
	go write("hello, GO")
	write("programing in GO")
}

func write(text string) {
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
