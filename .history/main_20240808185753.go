package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	const remainingTickets = 50

	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	fmt.Println("Enter your first name please:")
	var firstName string
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name please:")
	fmt.Scanln()
	var lastName string
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address please:")
	fmt.Scanln()
	var email string
	fmt.Scan(&email)
 
	fmt.Println("Enter the number of tickets you want to book:")
	fmt.Scanln()
	var userTickets int
	fmt.Scan(&userTickets)
	
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, , userTicketsemail)

}
