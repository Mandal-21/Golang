package main

import (
	"booking-app/helper"
	"fmt"
)

// Global variables
// constant variable
var conferenceName = "Go conference"

const conferenceTickets int = 50

var remainingTickets uint = 50

// var booking [50]string // array of fixed length strings. Its limit is 50
// var booking = []string{} // In slices we dont have to worry about length
// var booking = make([]map[string]string, 0)
var booking = make([]userData, 0)

type userData struct {
	firstName   string
	lastName    string
	email       string
	userTickets uint
}

func main() {

	fmt.Println("points to the value", conferenceName)
	fmt.Println("points to the memory address of conferenceName", &conferenceName)

	// remainingTickets = -100 wont work since uint is always positive

	greetUsers()

	// add first name and last name to booking

	// booking[0] = firstName + " " + lastName
	// fmt.Printf("The whole array %v\n", booking)
	// fmt.Printf("The first value %v\n", booking[0])
	// fmt.Printf("The first value %T\n", booking)
	// fmt.Printf("Array length: %v\n", len(booking))

	for {

		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			// calculate remaining tickets
			bookTicket(firstName, lastName, userTickets, email)

			firstNames := getFirstNames()
			fmt.Printf("First names from the booking are: %v\n", firstNames)

			noTicketsRemaining := remainingTickets == 0
			if noTicketsRemaining {
				// exit for loop
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Please enter valid first and last name")
			}
			if !isValidEmail {
				fmt.Println("Please enter valid email")
			}
			if !isValidTicketNumber {
				fmt.Println("Please enter valid ticket number")
			}
		}

	}

	// Switch Statements
	city := "London"
	switch city {
	case "London":
		fmt.Println("London is the capital city of England.")
	case "Paris":
		fmt.Println("Paris is the capital city of France.")
	case "New York":
		fmt.Println("New York is the capital city of USA.")
	case "Mexico", "Berlin":
		fmt.Println("Berlin is the capital city of Germany.")
	default:
		fmt.Println("No Valid city selected")
	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v our conference booking\n", conferenceName)
	fmt.Println("We have total", conferenceTickets, "tickets and", remainingTickets, "tickets remaining")
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, book := range booking {
		firstNames = append(firstNames, book.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask user name
	fmt.Println("Please enter first name")
	fmt.Scan(&firstName)
	fmt.Println("Please enter last name")
	fmt.Scan(&lastName)
	fmt.Println("Please enter email")
	fmt.Scan(&email)
	fmt.Println("Please enter number of tickets")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(firstName string, lastName string, userTickets uint, email string) []userData {
	remainingTickets = remainingTickets - userTickets
	// booking = append(booking, firstName+" "+lastName) //slice
	// make key value pairs
	// var userData = make(map[string]string)
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	// struct
	userData := userData{
		firstName:   firstName,
		lastName:    lastName,
		email:       email,
		userTickets: userTickets,
	}

	booking = append(booking, userData)
	fmt.Printf("List of bookings: %v\n", booking)

	fmt.Printf("User %v booked %v tickets. You will get confirmation at %v \n", firstName, userTickets, email)
	fmt.Printf("%v tickets remaining for conference %v \n", remainingTickets, conferenceName)
	return booking
}
