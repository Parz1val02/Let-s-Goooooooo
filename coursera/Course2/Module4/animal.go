package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}
type Cow struct {
	name       string
	food       string
	locomotion string
	noise      string
}
type Bird struct {
	name       string
	food       string
	locomotion string
	noise      string
}
type Snake struct {
	name       string
	food       string
	locomotion string
	noise      string
}

func (c Cow) Eat() {
	fmt.Println("Food eaten by " + c.name + "> " + c.food)
}

func (c Cow) Move() {
	fmt.Println("Locomotion method of " + c.name + "> " + c.locomotion)
}

func (c Cow) Speak() {
	fmt.Println("Spoken sound by i" + c.name + "> " + c.noise)
}

func (s Snake) Eat() {
	fmt.Println("Food eaten by " + s.name + "> " + s.food)
}

func (s Snake) Move() {
	fmt.Println("Locomotion method of " + s.name + "> " + s.locomotion)
}

func (s Snake) Speak() {
	fmt.Println("Spoken sound by " + s.name + "> " + s.noise)
}

func (b Bird) Eat() {
	fmt.Println("Food eaten by i" + b.name + "> " + b.food)
}

func (b Bird) Move() {
	fmt.Println("Locomotion method of " + b.name + "> " + b.locomotion)
}

func (b Bird) Speak() {
	fmt.Println("Spoken sound by " + b.name + "> " + b.noise)
}

func main() {
	cows := make([]Cow, 0)
	birds := make([]Bird, 0)
	snakes := make([]Snake, 0)
	for {
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("> ")
		request, _ := reader.ReadString('\n')
		request = strings.TrimSpace(request)

		input := strings.Fields(request)
		if len(input) < 3 {
			fmt.Println("Not enough arguments. Try again")
			continue
		}
		name := input[1]
		switch strings.ToLower(input[0]) {
		case "newanimal":
			animal := input[2]
			switch strings.ToLower(animal) {
			case "cow":
				c := Cow{name, "grass", "walk", "moo"}
				cows = append(cows, c)
			case "snake":
				s := Snake{name, "mice", "slither", "hsss"}
				snakes = append(snakes, s)
			case "bird":
				b := Bird{name, "worms", "sly", "peep"}
				birds = append(birds, b)
			default:
				fmt.Println("Not a valid animal")
				continue
			}
			fmt.Println("Created it!")
		case "query":
			action := input[2]
			var cowOpt Cow
			var snakeOpt Snake
			var birdOpt Bird
			for _, i := range cows {
				if name == i.name {
					cowOpt = i
					break
				}
			}
			for _, i := range snakes {
				if name == i.name {
					snakeOpt = i
					break
				}
			}
			for _, i := range birds {
				if name == i.name {
					birdOpt = i
					break
				}
			}
			if cowOpt.name == "" && snakeOpt.name == "" && birdOpt.name == "" {
				fmt.Println("Name not found")
				continue
			}
			switch strings.ToLower(action) {
			case "eat":
				if cowOpt.name != "" {
					cowOpt.Eat()
				} else if snakeOpt.name != "" {
					snakeOpt.Eat()
				} else {
					birdOpt.Eat()
				}
			case "move":
				if cowOpt.name != "" {
					cowOpt.Move()
				} else if snakeOpt.name != "" {
					snakeOpt.Move()
				} else {
					birdOpt.Move()
				}
			case "speak":
				if cowOpt.name != "" {
					cowOpt.Speak()
				} else if snakeOpt.name != "" {
					snakeOpt.Speak()
				} else {
					birdOpt.Speak()
				}
			default:
				fmt.Println("Not a valid action")
				continue
			}
		default:
			fmt.Println("Not a valid argument")
			continue
		}
	}
}
