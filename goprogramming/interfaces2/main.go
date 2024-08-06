package main

import (
	"errors"
	"fmt"
	"os"
)

type Employee struct {
	Id        int
	FirstName string
	LastName  string
}

type Developer struct {
	Individual        Employee
	HourlyRate        float64
	HoursWorkedInYear float64
	Review            map[string]interface{}
}

type Manager struct {
	Individual     Employee
	Salary         float64
	CommissionRate float64
}

type Payer interface {
	Pay() (string, float64)
}

func (dev Developer) Pay() (string, float64) {
	name := dev.Individual.FirstName + " " + dev.Individual.LastName
	return name, dev.HourlyRate * dev.HoursWorkedInYear
}

func (dev Developer) Rating() error {
	total := 0
	for _, i := range dev.Review {
		rating, err := compare(i)
		if err != nil {
			return err
		}
		total += rating
	}
	averageRating := float64(total) / float64(len(dev.Review))
	fmt.Printf("%s got a review rating of %.2f\n", dev.Individual.FirstName+" "+dev.Individual.LastName, averageRating)
	return nil
}

func (man Manager) Pay() (string, float64) {
	name := man.Individual.FirstName + " " + man.Individual.LastName
	return name, man.Salary + (man.Salary * man.CommissionRate)
}

func payDetails(p Payer) {
	name, pay := p.Pay()
	fmt.Printf("Name: %s, Pay: %0.2f\n", name, pay)
}

func compare(i interface{}) (int, error) {
	switch v := i.(type) {
	case int:
		return v, nil
	case string:
		rating, err := convert(v)
		if err != nil {
			return 0, err
		}
		return rating, nil
	default:
		return 0, errors.New("unknown type")
	}
}

func convert(s string) (int, error) {
	switch s {
	case "Excellent":
		return 5, nil
	case "Good":
		return 4, nil
	case "Fair":
		return 3, nil
	case "Poor":
		return 2, nil
	case "Unsatisfactory":
		return 1, nil
	default:
		return 0, errors.New("invalid rating: " + s)
	}
}

func main() {
	employeeReview := make(map[string]interface{})
	employeeReview["WorkQuality"] = 5
	employeeReview["TeamWork"] = 2
	employeeReview["Communication"] = "Poor"
	employeeReview["Problem-solving"] = 4
	employeeReview["Dependability"] = "Unsatisfactory"
	d := Developer{Individual: Employee{Id: 1, FirstName: "Eric", LastName: "Davis"}, HourlyRate: 35, HoursWorkedInYear: 2400, Review: employeeReview}
	m := Manager{Individual: Employee{Id: 2, FirstName: "Mr.", LastName: "Boss"}, Salary: 150000, CommissionRate: .07}
	err := d.Rating()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	payDetails(d)
	payDetails(m)
}
