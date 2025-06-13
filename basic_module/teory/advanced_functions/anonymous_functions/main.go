package main

import "fmt"

func main() {

	// way 1, no parameters
	func() {
		fmt.Println("hello, go!")
	}()

	//way 2, with parameters
	func(text string) {
		fmt.Println(text)
	}("anonymous function with parameter")

	//way 3, with return
	returned := func(text string) string {
		return fmt.Sprintf(" recived -> %s", text)
	}("anonymous function with parameter")

	fmt.Println(returned)

}
