package main

import (
	"fmt"
	"strings"
)

//-----package level variables declaration----
var conferenceName = "Go Conference"  
const conferenceTickets int = 50
var remainingTickets uint = 50
var bookings []string   //this declaration is slices, an abstraction of arrays with dynamic size. No need to specify size here.


func main(){
	for {
		greetUser()
		
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if  isValidName && isValidEmail && isValidTicketNumber{
			
			bookTicket(firstName, lastName, email, userTickets)
			firstNames := getFirstNames()
			fmt.Println("The first names of the bookings are", firstNames)
			if remainingTickets == 0 {
				fmt.Println("Our conference is full, try again next year.")
				break
			}
		} else {
			if !isValidName{
				fmt.Println("Name is invalid.")
			}
			if !isValidEmail {
				fmt.Println("Email is invalid.")
			}
			if !isValidTicketNumber {
				fmt.Println("Ticket number is invalid.")
			}
		}
	}
}

func greetUser(){
	fmt.Printf("Welcome to our %v booking application.\n", conferenceName)
	fmt.Println("We have a total of", conferenceTickets,"tickets and", remainingTickets,"tickets are still available.")
	fmt.Println("Get you bookings now!!")
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	
	//ask for user input
	fmt.Println("Enter user first name:")
	fmt.Scan(&firstName)  //fmt.Scan() needs a pointer reference to take input. It cannot directly take input into the variable.
	fmt.Println("Enter user last name:")
	fmt.Scan(&lastName)
	fmt.Println("Enter user email address:")
	fmt.Scan(&email)
	fmt.Println("Enter the number of tickets to be booked:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) > 2 && len(lastName)>2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}

func bookTicket(firstName string, lastName string, email string, userTickets uint){
	remainingTickets = remainingTickets - userTickets
	// bookings[0] = firstName + " " + lastName
	bookings = append(bookings, firstName+ " "+ lastName)
	fmt.Println("Thank You", firstName, lastName,"for booking", userTickets,"tickets. You will receive a confirmation email at",email)
	fmt.Println(remainingTickets,"tickets remainaing for",conferenceName)
}

func getFirstNames() []string{
	firstNames := []string{}
	for _,booking := range bookings{
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}