package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	const remainingTickets = 50

	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	fmt.Println("Enter your name please:")
	var userName string
	fmt.Scan(&userName)
 
	fmt.Println("Enter the number of tickets you want to book")
	var userTickets int
	fmt.Scan(&userTickets)
	
	
	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)

}
