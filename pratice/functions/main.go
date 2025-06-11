package main

import (
	"fmt"
)

func main() {
	fmt.Println(itsAnagram("amor", "romo"))
}

func sumPairs(a []int16) {
	sum := int16(0)

	for _, v := range a {
		if v%2 == 0 {
			sum += v
		}
	}

	fmt.Println("the sum pairs it's", sum)
}

func fatorial(n int8) int8 {
	var result int8 = 1
	for i := n; i > 1; i-- {
		result = result * i
	}

	return result

}

// revisar essa questão
/* func itsAnagram(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	// Mapa para contar as letras da primeira string
	letras := make(map[rune]int)

	for _, char := range a {
		letras[char]++
	}

	// Subtrai as contagens com base na segunda string
	for _, char := range b {
		letras[char]--
		if letras[char] < 0 {
			return false
		}
	}

	return true
} */

func itsAnagram(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	return true
}
