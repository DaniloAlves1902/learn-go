package main

import (
	"fmt"
	//"time"
)

func main() {

	// way 1, similar to while
	i := 0

	for i < 3 {
		i++
		fmt.Printf("i agora vale: %d \n", i)
		//time.Sleep(time.Second)
	}

	fmt.Println("-------------------------------")

	// way 2, similar to "for" conventional
	for j := 0; j < 3; j++ {
		fmt.Printf("j agora vale: %d \n", j)
		//time.Sleep(time.Second)
	}

	fmt.Println("-------------------------------")

	names := [3]string{"a", "b", "c"}

	for i := range names {
		fmt.Println(names[i])
	}

	fmt.Println("-------------------------------")

	for i, name := range names {
		fmt.Printf("names[%d] = %s\n", i, name)
	}

	fmt.Println("-------------------------------")

	for _, name := range names {
		fmt.Printf("%s\n", name)
	}

	fmt.Println("-------------------------------")

	for i, phrase := range "I LOVE SARAH JENNYFER" {
		fmt.Println(i, string(phrase))
	}

	fmt.Println("-------------------------------")

	users := map[string]string{
		"name":     "Sarah Jennyfer",
		"position": "graphic design",
	}

	for key, value := range users {
		fmt.Printf("%s %s\n", key, value)
	}

	/*
		infinty loop
		
		for {
			fmt.Println("hi, go")
		}
	*/

}
