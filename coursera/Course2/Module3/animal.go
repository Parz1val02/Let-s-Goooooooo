package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (animal Animal) Eat() {
	fmt.Println("Food eaten> " + animal.food)
}

func (animal Animal) Move() {
	fmt.Println("Locomotion method> " + animal.locomotion)
}

func (animal Animal) Speak() {
	fmt.Println("Spoken sound> " + animal.noise)
}

func main() {
	var animal Animal
	for {
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Type a request (e.g. animal info)> ")
		request, _ := reader.ReadString('\n')
		request = strings.TrimSpace(request)

		input := strings.Fields(request)
		switch strings.ToLower(input[0]) {
		case "cow":
			animal.food = "grass"
			animal.locomotion = "walk"
			animal.noise = "moo"
		case "bird":
			animal.food = "worms"
			animal.locomotion = "fly"
			animal.noise = "peep"
		case "snake":
			animal.food = "mice"
			animal.locomotion = "slither"
			animal.noise = "hsss"
		default:
			fmt.Println("Not a valid animal")
			continue
		}
		fmt.Printf("Animal> %v\n", input[0])
		switch strings.ToLower(input[1]) {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
		default:
			fmt.Println("Not a valid action")
		}
	}
}
