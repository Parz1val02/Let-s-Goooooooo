package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
	//"strconv"
	"strings"
)

var conferenceName string = "Go Conference"
const conferenceTickets int = 50
var remainingTickets uint = 50
//var bookings [50]string Arrays: fixed size
//Slices: fixed size 
// List of maps -> var bookings = make([]map[string]string, 0)
var bookings = make([]UserData, 0)

type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	//fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)
	greetUsers()

//	for {
		firstName, lastName, userTickets, email := helper.GetUserInput()

		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber{
			bookTicket(userTickets, firstName, lastName, email)
			//Concurrency
		wg.Add(1)
			go sendTicket(uint(userTickets), firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Printf("Our conference is booked out. Come back next year\n")
//				break
			}
		} else {
			if !isValidName{
				fmt.Println("First or last name entered is too short")
			}
			if !isValidEmail{
				fmt.Println("Email address entered does not contain @ sign")
			}
			if !isValidTicketNumber{
				fmt.Println("Number of tickets entered is invalid")
			}
		}
	//}
	//bookings[0] = firstName + " " + lastName
	//fmt.Printf("The whole slice: %v\n", bookings)
	//fmt.Printf("The first value: %v\n", bookings[0])
	//fmt.Printf("Slice type: %T\n", bookings)
	//fmt.Printf("Slice length: %v\n", len(bookings))
	wg.Wait()
}
func bookTicket(userTickets int, firstName string, lastName string, email string){
	remainingTickets -= uint(userTickets)
	//Create a map for a user
	//var userData = make(map[string]string)
	//userData["firstName"] = firstName
	//userData["lastName"] = lastName
	//userData["email"] = email
	//userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	var userData = UserData {
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: uint(userTickets),
	}

	bookings = append(bookings, userData) 
	fmt.Printf("List of bookings: %v\n", bookings)

	fmt.Printf("Thank you %s %s for booking %d tickets. You will receive a confirmation email at %s\n",firstName,lastName,userTickets,email)
	fmt.Printf("%v tickets remaining for %v\n",remainingTickets,conferenceName)
}
 func getFirstNames() []string{
	firstNames := []string{}
	for _, booking := range bookings{
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}
func greetUsers(){
	fmt.Printf("Welcome to our %v booking application!\n", conferenceName)
	fmt.Println("We have a total of",conferenceTickets,"tickets and",remainingTickets,"are still available")
	fmt.Println("Get your tickets here to attend")
}
func validateUserInput(firstName string, lastName string, email string, userTickets int) (bool, bool, bool){
	var isValidName bool = len(firstName) >= 2 && len(lastName) >= 2
	var isValidEmail bool = strings.Contains(email, "@")
	var isValidTicketNumber bool = userTickets >0 && userTickets <= int(remainingTickets)
	return isValidName, isValidEmail, isValidTicketNumber
}
func sendTicket(userTickets uint, firstName string, lastName string, email string){
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v\n", userTickets, firstName, lastName)
	fmt.Println("###############################")
	fmt.Printf("Sending ticket:\n%v\nto email address %v\n", ticket, email)
	fmt.Println("###############################")
	wg.Done()
}
