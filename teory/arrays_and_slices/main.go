package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("---------------arrays---------------")
	var arr1 [5]int
	arr1[0] = 3

	fmt.Println(arr1)

	arr2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr2)

	arr3 := [...]int{1, 2, 3, 4}
	fmt.Println(arr3)

	fmt.Println(reflect.TypeOf(arr3))

	fmt.Println("---------------slices---------------")
	slice := []int{4, 6, 5, 7, 3}

	fmt.Println(slice)
	fmt.Println(reflect.TypeOf(slice))

	slice = append(slice, 12)
	fmt.Println(slice)

	slice2 := arr2[1:3]
	fmt.Println(slice2)

	fmt.Println("---------------internal arrays---------------")
	slice3 := make([]int, 5, 6)
	fmt.Println("initial size:", len(slice3))
	fmt.Println("total capacity:", cap(slice3)) // capacity

	slice3 = append(slice3, 1)
	slice3 = append(slice3, 2)
	fmt.Println("initial size:", len(slice3))
	fmt.Println("total capacity:", cap(slice3)) // capacity

}
