package main

import "fmt"

func main() {

	fmt.Println("----map----")

	user1 := map[string]string{
		"fName": "danilo",
		"lName": "F. Alves",
	}

	fmt.Println(user1["fName"])
	fmt.Println(user1)

	fmt.Println("----aligned map----")
	// map within a map

	user2 := map[string]map[string]string{
		"name": {
			"first": "Sarah",
			"last":  "J. V. Alves",
		},

		"course": {
			"name":   "IT",
			"school": "IFPB",
		},
	}

	fmt.Println()

	// learn later

	/* for category, data := range user2 {
		fmt.Println(category)
		for key, value := range data {
			fmt.Printf("- %s : %s", key, value)
			fmt.Println()
		}
	} */

	fmt.Println(user2["name"]["first"])
	fmt.Println(user2["course"]["name"])

	fmt.Println(user2)

	stock := map[string]map[string]int{
		"tech": {
			"keyboard":            10,
			"mouse":               8,
			"xbox series s":       3,
			"playstation 5":       6,
			"video card rtx 5090": 2,
			"video card rtx 4070": 3,
		},

		"clothes": {
			"t-shift violence is cool, thanks video games": 15,
			"t-shirt white": 20,
			"t-shirt black": 18,
			"t-shirt red":   7,
		},

		"mugs": {
			"java mug":    26,
			"GTA 6 mug":   10,
			"MLP 4th mug": 32,
		},
	}

	fmt.Println(stock["tech"]["xbox series s"])

	// learn later

	/* for category, products := range stock {
		fmt.Println("Categoria:", category)
		for name, quantity := range products {
			fmt.Printf("  - %s: %d unidades\n", name, quantity)
		}
		fmt.Println()
	} */
}
