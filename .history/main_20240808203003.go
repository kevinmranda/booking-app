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

		if userTickets <= remainingTickets {
			
		} else if userTickets > remainingTickets {
			fmt.Printf("There are %v tickets remaining. Try again correctly!\n", remainingTickets)
			continue
		}


	}
}
