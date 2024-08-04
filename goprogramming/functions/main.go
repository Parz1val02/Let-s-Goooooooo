package main

import "fmt"

type Employee struct {
	Id        int
	FirstName string
	LastName  string
}

type Developer struct {
	Individual Employee
	HourlyRate int
	WorkWeek   [7]int
}

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func (dev *Developer) LogHours(weekday Weekday, num int) {
	dev.WorkWeek[weekday] = num
}

func (dev *Developer) HoursWorked() int {
	total := 0
	for _, v := range dev.WorkWeek {
		total += v
	}
	return total
}

func (dev *Developer) PayDay() (int, bool) {
	worked := dev.HoursWorked()
	var payment int
	if worked <= 40 {
		payment = worked * dev.HourlyRate
		return payment, false
	} else {
		payment = (40 * dev.HourlyRate) + (worked-40)*2*dev.HourlyRate
		return payment, true
	}
}

func (dev *Developer) PayDetails() {
	for i, v := range dev.WorkWeek {
		switch i {
		case 0:
			fmt.Println("Sunday hours: ", v)
		case 1:
			fmt.Println("Monday hours: ", v)
		case 2:
			fmt.Println("Tuesday hours: ", v)
		case 3:
			fmt.Println("Wednesday hours: ", v)
		case 4:
			fmt.Println("Thursday hours: ", v)
		case 5:
			fmt.Println("Friday hours: ", v)
		case 6:
			fmt.Println("Saturday hours: ", v)
		}
	}

	fmt.Printf("\nHours worked this week:  %d\n", dev.HoursWorked())
	pay, overtime := dev.PayDay()
	fmt.Println("Pay for the week: $", pay)
	fmt.Println("Is this overtime pay: ", overtime)
}

func main() {
	developer := Developer{Individual: Employee{Id: 1, FirstName: "Tony", LastName: "Stark"}, HourlyRate: 10}

	n := nonLoggedHours()
	fmt.Printf("Tracking hours worked thus far: %d\n", n(2))
	fmt.Printf("Tracking hours worked thus far: %d\n", n(3))
	fmt.Printf("Tracking hours worked thus far: %d\n", n(5))
	fmt.Println()

	developer.LogHours(Monday, 8)
	developer.LogHours(Tuesday, 10)
	developer.LogHours(Wednesday, 10)
	developer.LogHours(Thursday, 10)
	developer.LogHours(Friday, 6)
	developer.LogHours(Saturday, 8)

	developer.PayDetails()
}

func nonLoggedHours() func(int) int {
	hours := 0
	return func(i int) int {
		hours += i
		return hours
	}
}
