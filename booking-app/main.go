package main

import (
	"fmt"
)

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("conferenceName is %T, conferenceTickets is %T, remainingTickets is %T.\n", conferenceName, conferenceTickets, remainingTickets)
	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still remaining.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend ")

	var userName string
	var userTickets int

	userName = "Athul"
	userTickets = 2

	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)
}
