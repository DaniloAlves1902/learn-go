package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	channel := multiplexer(write("Hello, Go"), write("programing in Go"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-channel);
	}

	

}

func multiplexer(inputChan1, inputChan2 <-chan string) <-chan string {
	outputChan := make(chan string)

	go func() {
		for {
			select {
			case message := <-inputChan1:
				outputChan <- message
			case message := <-inputChan2:
				outputChan <- message
			}
		}
	}()

	return outputChan
}

func write(text string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			channel <- fmt.Sprintf("reciver value: %s", text)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return channel
}
