package main

import (
	"fmt"
	"os"
)

type Saiyan struct {
	Name   string
	Power  int
	Father *Saiyan
}

func (s *Saiyan) Super2() {
	s.Power += 2000
}

type person struct {
	Name string
}

func (p *person) Introduce() {
	fmt.Printf("Hi, I'm %s\n", p.Name)
}

type Saiyan2 struct {
	*person
	Power int
}

// Constructor
func NewSaiyan(name string, power int) *Saiyan {
	return &Saiyan{
		Name:  name,
		Power: power,
	}
}

func main() {
	//Command line arguments
	if len(os.Args) != 2 {
		os.Exit(1)
	}
	fmt.Println("It's over", os.Args[1])

	//Assign multiples variables
	name, power := "Goku", 9000
	fmt.Printf("%s's power is over %d\n", name, power)

	goku := &Saiyan{Name: "Goku"}
	goku.Power = 9000
	Super(goku)
	goku.Super2()
	fmt.Println(goku.Power)

	vegetta := NewSaiyan("Vegetta", 10000)
	fmt.Printf("%s's power is over %d\n", vegetta.Name, vegetta.Power)

	gohan := new(Saiyan)
	gohan.Name = "Gohan"
	gohan.Power = 20000
	gohan.Father = &Saiyan{
		Name:   "Goku",
		Power:  1000000000,
		Father: nil,
	}
	fmt.Printf("%s's power is over %d and his dad is %s\n", gohan.Name, gohan.Power, gohan.Father.Name)

	goku2 := &Saiyan2{
		person: &person{"Goku"},
		Power:  90001,
	}
	goku2.Introduce()
}

func Super(s *Saiyan) {
	s.Power += 1000
}
