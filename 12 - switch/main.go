package main

import "fmt"

func weekDay(num int) string {
	switch num {
	case 1:
		return "Sunday"
	case 2:
		return "Monday"
	case 3:
		return "Tuesday"
	case 4:
		return "Wednesday"
	case 5:
		return "Thursday"
	case 6:
		return "Friday"
	case 7:
		return "Saturday"
	default:
		return "Invalid"
	}
}

func weekDay2(num int) string {
	var weekDay string

	switch {
	case num == 1:
		weekDay = "Sunday"
	case num == 2:
		weekDay = "Monday"
	case num == 3:
		weekDay = "Tuesday"
	case num == 4:
		weekDay = "Wednesday"
	case num == 5:
		weekDay = "Thursday"
	case num == 6:
		weekDay = "Friday"
	case num == 7:
		weekDay = "Saturday"
	default:
		weekDay = "Invalid"
	}

	return weekDay
}

func main() {
	day := weekDay(200)
	fmt.Println(day)

	fmt.Println("-----------")
	
	day2 := weekDay2(1)
	fmt.Println(day2)
}