package main

import (
	"fmt"
	"strings"
)

func main() {
	// constant variable
	conferenceName := "Go conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	// var booking [50]string // array of fixed length strings. Its limit is 50
	booking := []string{} // In slices we dont have to worry about length

	fmt.Println(conferenceName)
	fmt.Println(&conferenceName) // points to the memory address of conferenceName

	// remainingTickets = -100 wont work since uint is always positive

	fmt.Printf("conferenceTickets: %T, conferenceName: %T, remainingTickets: %T\n", conferenceName, remainingTickets, conferenceTickets)

	fmt.Printf("Welcome to %v our conference booking\n", conferenceName)
	fmt.Println("We have total", conferenceTickets, "tickets and", remainingTickets, "tickets remaining")
	fmt.Println("Get your tickets here to attend")

	// add first name and last name to booking

	// booking[0] = firstName + " " + lastName
	// fmt.Printf("The whole array %v\n", booking)
	// fmt.Printf("The first value %v\n", booking[0])
	// fmt.Printf("The first value %T\n", booking)
	// fmt.Printf("Array length: %v\n", len(booking))

	for {

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

		if userTickets <= remainingTickets {
			// calculate remaining tickets
			remainingTickets = remainingTickets - userTickets

			booking = append(booking, firstName+" "+lastName) //slice
			fmt.Printf("The whole slice: %v\n", booking)
			fmt.Printf("The first value %v\n", booking[0])
			fmt.Printf("Slice type: %T\n", booking)
			fmt.Printf("Slice length: %v\n", len(booking))

			fmt.Printf("User %v booked %v tickets. You will get confirmation at %v \n", firstName, userTickets, email)
			fmt.Printf("%v tickets remaining for conference %v \n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, book := range booking {
				var names = strings.Fields(book)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("First names from the booking are: %v\n", firstNames)

			noTicketsRemaining := remainingTickets == 0
			if noTicketsRemaining {
				// exit for loop
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			fmt.Println("Sorry, we dont have enough tickets")
		}

	}
}
