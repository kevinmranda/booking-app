package main

import (
	"fmt"
	"strings"
)

const conferenceTickets = 50
var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings []string
var firstName, lastName, email string
var userTickets uint

func main() {

	greetUsers()

	for {

		firstName, lastName, email, userTickets := getUserInputs()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput()

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(remainingTickets, userTickets, bookings, firstName, lastName, email, conferenceName)

			firstNames := getFirstNames()

			fmt.Printf("These are all our bookings: %v\n", firstNames)
			fmt.Println()

			if remainingTickets == 0 {
				fmt.Printf("Sorry! %v is booked out.", conferenceName)
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or Last name you entered is too short.")
			}
			if !isValidEmail {
				fmt.Println("The email address you entered doesn't contain @ sign.")
			}
			if !isValidTicketNumber {
				fmt.Println("The number of tickets you entered is invalid.")
			}
		}
	}

}

func greetUsers() {
	fmt.Printf(
		"Welcome to our %v booking application\n",
		conferenceName,
	)
	fmt.Printf(
		"We have a total of %v tickets and %v are still available\n",
		conferenceTickets,
		remainingTickets,
	)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}

	return firstNames
}

func validateUserInput() (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}

func getUserInputs() (string, string, string, uint) {
	fmt.Println("Enter your first name please:")
	var firstName string
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name please:")
	var lastName string
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address please:")
	var email string
	fmt.Scan(&email)

	fmt.Println("Enter the number of tickets you want to book:")
	var userTickets uint
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(conName string) []string {
	remainingTickets -= userTickets

	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining are %v\n", conName, remainingTickets)

	return bookings
}
