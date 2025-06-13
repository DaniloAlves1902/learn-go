package main

import "fmt"

func generic(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generic("string")
	generic(1)
	generic(true)

	genericMap := map[interface{}]interface{}{
		"string": 1,
		true:     float64(10.98),
	}

	fmt.Println(genericMap)
}
