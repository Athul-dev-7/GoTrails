package main

import (
	"fmt"
	"strings"
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

	for {
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

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 1 && userTickets <= uint(remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - int(userTickets)
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thankyou %v %v for booking %v tickets. You will receive a confirmation mail at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v remaining tickets for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings {
				names := strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("The first names of bookings are : %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our booking is closed for now. Stay updated for later announcement")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Your firstname or lastname entered is too short.")
			}
			if !isValidEmail {
				fmt.Println("You mail doesn't contain '@' sign")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid")
			}
		}

	}

}
