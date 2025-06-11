package main

import "fmt"

func main() {
	var variable1 string = "variable declared"

	variable2 := "variable inferred"

	fmt.Println(variable1)
	fmt.Println(variable2)

	var (
		variable3 string = "test 1"
		variable4 string = "test 2"
	)

	fmt.Println(variable3, variable4)

	variable5, variable6 := "test 3", "test 4"

	fmt.Println(variable5, variable6)

	variable5, variable6 = variable6, variable5

	fmt.Println(variable5, variable6)

    //exercise

}
