package main

import (
	"fmt"
	"sync"
	"time"
	"os"
	"date"
	"log"
)

const conferenceTickets = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

// Struct UserData
type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

//interface
type IAccount interface{
	withdraw()
}

type Account struct{
	amount float64
	date date.Time
}

func ()

// Goroutines
var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {

		bookTicket(userTickets, firstName, lastName, email)

		wg.Add(1)
		//concurrency
		go sendTicket(userTickets, firstName, lastName, email)

		firstNames := getFirstNames()

		fmt.Printf("These are all our bookings: %v\n", firstNames)
		fmt.Println()

		if remainingTickets == 0 {
			fmt.Printf("Sorry! %v is booked out.", conferenceName)
			// break
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
	wg.Wait()
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
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
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

func bookTicket(userTickets uint, firstName, lastName, email string) {
	remainingTickets -= userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of Bookings is: %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining are %v\n", conferenceName, remainingTickets)
}

func sendTicket(userTickets uint, firstName, lastName, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("----------------------")
	fmt.Printf("sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("----------------------")
	wg.Done()
}

func openFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %w", filename, err)
	}
	defer f.Close()
	return nil
}
