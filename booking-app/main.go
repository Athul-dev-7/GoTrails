package main

import (
	"fmt"
)

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50
	var bookings []string

	fmt.Printf("conferenceName is %T, conferenceTickets is %T, remainingTickets is %T.\n", conferenceName, conferenceTickets, remainingTickets)
	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still remaining.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend ")

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// Ask user for their details
	fmt.Println("Enter your first name")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address")
	fmt.Scan(&email)

	fmt.Println("Enter the number of tickets")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - int(userTickets)
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thankyou %v %v for booking %v tickets. You will receive a confirmation mail at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v remaining tickets for %v\n", remainingTickets, conferenceName)

	fmt.Printf("These are all our bookings :%v\n", bookings)
}
