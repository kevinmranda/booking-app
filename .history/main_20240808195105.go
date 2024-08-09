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

	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

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

		remainingTickets -= userTickets

		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining are %v\n", conferenceName, remainingTickets)

		firstNames := []string{}
		for index, booking := range bookings {
			var names = strings.Fields(booking)
			append(firstNames, )
		}

		fmt.Printf("These are all our bookings: %v\n", bookings)
		fmt.Println()
	}
}
