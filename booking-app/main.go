package main

import "fmt"

func main() {
	// constant variable
	conferenceName := "Go conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	fmt.Println(conferenceName)
	fmt.Println(&conferenceName) // points to the memory address of conferenceName

	// remainingTickets = -100 wont work since uint is always positive

	fmt.Printf("conferenceTickets: %T, conferenceName: %T, remainingTickets: %T\n", conferenceName, remainingTickets, conferenceTickets)

	fmt.Printf("Welcome to %v our conference booking\n", conferenceName)
	fmt.Println("We have total", conferenceTickets, "tickets and", remainingTickets, "tickets remaining")
	fmt.Println("Get your tickets here to attend")

	var firstName string
	var lastName string
	var email string
	var userTickets int
	// ask user name
	fmt.Println("Please enter first name")
	fmt.Scan(&firstName)
	fmt.Scan(&lastName)
	fmt.Scan(&email)
	fmt.Scan(&userTickets)

	
	userTickets = 2
	fmt.Printf("User %v booked %v tickets. You will get confirmation at %v \n", firstName, userTickets, email)

}