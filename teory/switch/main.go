package main

import "fmt"

//way 1
func dayOffWeek(number int) string {
	switch number {
	case 1:
		return "Sunday"
	case 2:
		return "monday"
	case 3:
		return "Tuesday"
	case 4:
		return "Wednesday"
	case 5:
		return "Thursday"
	case 6:
		return "Saturday"
	case 7:
		return "monday"
	default:
		return "invalid number, try again"
	}

}

//way 2
func dayOffWeek2(number int) string {
	switch {
	case number == 1:
		return "Sunday"
	case number == 2:
		return "monday"
	case number == 3:
		return "Tuesday"
	case number == 4:
		return "Wednesday"
	case number == 5:
		return "Thursday"
	case number == 6:
		return "Saturday"
	case number == 7:
		return "monday"
	default:
		return "invalid number, try again"
	}

}

//way 3
func dayOffWeek3(number int) string {
	var dayOffWeek string
	switch {
	case number == 1:
		dayOffWeek = "Sunday"
	case number == 2:
		dayOffWeek = "monday"
	case number == 3:
		dayOffWeek = "Tuesday"
	case number == 4:
		dayOffWeek = "Wednesday"
	case number == 5:
		dayOffWeek = "Thursday"
	case number == 6:
		dayOffWeek = "Saturday"
	case number == 7:
		dayOffWeek = "monday"
	default:
		dayOffWeek = "invalid number, try again"
	}

	return dayOffWeek

}

func main() {

	var day int

	fmt.Print("enter a number from 0 to 7: ")
	fmt.Scanln(&day)

	fmt.Println(dayOffWeek(day))
	fmt.Println("-------------")
	fmt.Println(dayOffWeek2(day))
	fmt.Println("-------------")
	fmt.Println(dayOffWeek3(day))
}
