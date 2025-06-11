package main

import (
	"errors"
	"fmt"
)

func main() {
	var number int32 = 2000

	fmt.Println(number)
	
	// não aceita negativos
	var number2 uint32 = 1234

	fmt.Println(number2)

	var numberFloat1 float32 = 15.32

	fmt.Println(numberFloat1)

	var boolean bool = true

	fmt.Println(boolean)

	var err error = errors.New("internal error")

	fmt.Println(err)
}