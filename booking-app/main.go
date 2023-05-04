package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

//-----package level variables declaration----
var conferenceName = "Go Conference"  
const conferenceTickets int = 50
var remainingTickets uint = 50
// var bookings = make([]map[string]string, 0)   //this declaration is slices, an abstraction of arrays with dynamic size. No need to specify size here.

var bookings = make([]UserData,0)
type UserData struct{
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}
func main(){

	greetUser()
	
	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if  isValidName && isValidEmail && isValidTicketNumber{
		
		bookTicket(firstName, lastName, email, userTickets)
		wg.Add(1)
		go sendTickets(userTickets, firstName, lastName, email)
		firstNames := getFirstNames()
		fmt.Println("The first names of the bookings are", firstNames)
		if remainingTickets == 0 {
			fmt.Println("Our conference is full, try again next year.")
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
	wg.Wait()

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


func bookTicket(firstName string, lastName string, email string, userTickets uint){
	remainingTickets = remainingTickets - userTickets
	// bookings[0] = firstName + " " + lastName
	// var userData = make(map[string]string)
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)
	var userData = UserData{
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}
	bookings = append(bookings, userData)
	fmt.Println("Thank You", firstName, lastName,"for booking", userTickets,"tickets. You will receive a confirmation email at",email)
	fmt.Println(remainingTickets,"tickets remainaing for",conferenceName)
	fmt.Println(bookings)
}

func getFirstNames() []string{
	firstNames := []string{}
	for _,booking := range bookings{
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func sendTickets(userTickets uint, firstName string, lastName string, email string){
	time.Sleep(50*time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("------------------------------")
	fmt.Printf("Sending ticket:\n %v \n to email address %v\n", ticket, email)
	fmt.Println("------------------------------")
	wg.Done()
}