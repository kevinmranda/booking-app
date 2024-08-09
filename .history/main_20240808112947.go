package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	const remainingTickets = 50

	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var  userName string
	fmt.Println("Enter you name here")
	fmt.Scan(userName)
	
	var userTickets int
	fmt.Println("Enter the number of of tickets")
	fmt.Scan(userName)
	// userName = "Kevin"
	// userTickets = 9

	fmt.Println(userName, userTickets)
	
}
