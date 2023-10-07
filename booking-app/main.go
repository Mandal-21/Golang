package main

import "fmt"

func main() {
	// constant variable
	conferenceName := "Go conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	// remainingTickets = -100 wont work since uint is always positive

	fmt.Printf("conferenceTickets: %T, conferenceName: %T, remainingTickets: %T\n", conferenceName, remainingTickets, conferenceTickets)

	fmt.Printf("Welcome to %v our conference booking\n", conferenceName)
	fmt.Println("We have total", conferenceTickets, "tickets and", remainingTickets, "tickets remaining")
	fmt.Println("Get your tickets here to attend")

	var userName string
	var userTickets int
	// ask user name
	userName = "tom"
	userTickets = 2
	fmt.Printf("User %v booked %v tickets\n", userName, userTickets)

}