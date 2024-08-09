package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	var bookings []string

	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	for {
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

		isValidName, is validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			remainingTickets -= userTickets

			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining are %v\n", conferenceName, remainingTickets)

			firstNames := getFirstNames(bookings)

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

func greetUsers(conName string, conTickets int, remTickets uint) {
	fmt.Printf(
		"Welcome to our %v booking application\n",
		conName,
	)
	fmt.Printf(
		"We have a total of %v tickets and %v are still available\n",
		conTickets,
		remTickets,
	)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}

	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}
