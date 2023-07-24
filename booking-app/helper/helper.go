package helper

import (
	"fmt"
)

func GetUserInput() (string, string, int, string){
	var firstName string
	var lastName string
	var userTickets int
	var email string
	fmt.Print("Enter your first name> ")
	fmt.Scanf("%s",&firstName)
	fmt.Print("Enter your last name> ")
	fmt.Scanf("%s",&lastName)
	fmt.Print("Enter your email> ")
	fmt.Scanf("%s",&email)
	fmt.Print("Enter number of tickets> ")
	fmt.Scanf("%d",&userTickets)
	return firstName, lastName, userTickets, email	
}
