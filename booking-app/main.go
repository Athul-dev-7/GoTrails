package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

const conferenceTickets = 50

var conferenceName = "Go Conference"
var remainingTickets = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUser()

	firstName, lastName, email, userTickets := getUserInputs()
	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {
		bookings := bookTickets(userTickets, firstName, lastName, email)

		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)

		firstNames := getFirstNames(bookings)
		fmt.Printf("The first names of bookings are : %v\n", firstNames)

		if remainingTickets == 0 {
			fmt.Println("Our booking is closed for now. Stay updated for later announcement")
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
	wg.Wait()

}

func greetUser() {
	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still remaining.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend ")
}

func getFirstNames(bookings []UserData) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInputs() (string, string, string, uint) {
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

	return firstName, lastName, email, userTickets
}

func bookTickets(userTickets uint, firstName string, lastName string, email string) []UserData {
	remainingTickets = remainingTickets - int(userTickets)

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}
	bookings = append(bookings, userData)
	fmt.Println("bookings are :", bookings)

	fmt.Printf("Thankyou %v %v for booking %v tickets. You will receive a confirmation mail at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v remaining tickets for %v\n", remainingTickets, conferenceName)
	return bookings
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v ", userTickets, firstName, lastName)
	fmt.Println("############################")
	fmt.Printf("Sending ticket : %v\nTo email address : %v\n", ticket, email)
	fmt.Println("############################")
	wg.Done()
}
