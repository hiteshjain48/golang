package main

import (
	"fmt"
	"strings"
)

func main(){
	conferenceName := "Go Conference"  //this syntax can be used if you don't want to explicitly specify the datatype or use var
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	fmt.Printf("Welcome to our %v booking application.\n", conferenceName)
	fmt.Println("We have a total of", conferenceTickets,"tickets and", remainingTickets,"tickets are still available.")
	fmt.Println("Get you bookings now!!")

	var bookings []string   //this declaration is slices, an abstraction of arrays with dynamic size. No need to specify size here.
	for {

		var firstName string
		var lastName string
		var email string
		var userTickets uint
		// var bookings [50]string
		
		//ask for user input
		fmt.Println("Enter user first name:")
		fmt.Scan(&firstName)  //fmt.Scan() needs a pointer reference to take input. It cannot directly take input into the variable.
		fmt.Println("Enter user last name:")
		fmt.Scan(&lastName)
		fmt.Println("Enter user email address:")
		fmt.Scan(&email)
		fmt.Println("Enter the number of tickets to be booked:")
		fmt.Scan(&userTickets)

		if userTickets <= remainingTickets {
			remainingTickets = remainingTickets - userTickets
			// bookings[0] = firstName + " " + lastName
			bookings = append(bookings, firstName+ " "+ lastName)
			fmt.Println("Thank You", firstName, lastName,"for booking", userTickets,"tickets. You will receive a confirmation email at",email)
			fmt.Println(remainingTickets,"tickets remainaing for",conferenceName)
	
			firstNames := []string{}
			for _,booking := range bookings{
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Println("The first names of the bookings are", firstNames)
			if remainingTickets == 0 {
				fmt.Println("Our conference is full, try again next year.")
				break
			}
		} else {
			fmt.Println("Only",remainingTickets,"remaining. You cannot book", userTickets,"tickets.")
			continue
		}
	}
}