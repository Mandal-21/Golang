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

	// calculate remaining tickets
	remainingTickets = remainingTickets - userTickets

	userTickets = 2
	fmt.Printf("User %v booked %v tickets. You will get confirmation at %v \n", firstName, userTickets, email)
	fmt.Printf("%v tickets remaining for conference %v \n", remainingTickets, conferenceName)

}
