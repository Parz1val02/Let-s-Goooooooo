package main

import (
	"fmt"
)

func main(){
	fmt.Println("Welcome to my quiz game!")
	//Implicit declaration
	var name string = "Rodro"
	//var int var1 = -1
	//var uint var2 = 2 //unsigned integer
	//var float64 var3 = 4.2 //define precision for float, also available float32
	//var ga bool = true
	fmt.Println(name)
	//Explicit declaration
	name2 := "Kipi"
	age := 21
	money := 33.3
	fmt.Printf("Hello %v, how are you doing\n", name2)
	fmt.Printf("You are %d years old and have %f soles in the bank\n", age, money)

	var ga string
	var age2 uint
	fmt.Print("Enter your name> ")
	fmt.Scanf("%s",&ga)
	fmt.Printf("Hello %s, welcome to the game\n",ga)
	fmt.Print("Enter your age> ")
	fmt.Scanf("%d",&age2)

	if age2 > 10 {
		fmt.Println("Tas basado mano")
	} else if age2 == 10 {
		fmt.Println("Puede ser puede ser")
	}else {
		fmt.Println("Sorry mano uwu")
	}
}
