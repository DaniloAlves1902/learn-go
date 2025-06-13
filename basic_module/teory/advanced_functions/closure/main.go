package main

import "fmt"

func closure() func() {
	text := "inside closure function"

	function := func() {
		fmt.Println(text)
	}

	return function
}

func main() {

	newFunction := closure()
	newFunction()

}
