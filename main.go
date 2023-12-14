package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go conference"
	const conferenceTickets uint8 = 50
	var remainingTickets uint8 = 50

	var bookings []string

	fmt.Printf("Welcome to %v our booking app\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are availabe.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {

		var firstName string
		var lastName string
		var email string
		var userTickets uint8

		fmt.Println()

		fmt.Println("Enter your first name")
		fmt.Scan(&firstName)

		fmt.Println()

		fmt.Println("Enter your last name")
		fmt.Scan(&lastName)

		fmt.Println()

		fmt.Println("Enter your email")
		fmt.Scan(&email)

		fmt.Println()

		setTicket(&userTickets, remainingTickets)

		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %v %v for bookin %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining\n", remainingTickets)

		firstNames := []string{}

		for _, booking := range bookings {
			names := strings.Fields(booking) // splits the string with white spaces as separator
			firstNames = append(firstNames, names[0])

		}

		fmt.Printf("Booking first names: %v\n", firstNames)

		if remainingTickets == 0 {
			fmt.Println("Our conference has no tickets left, until next year!")
			break
		}
	}
}

func setTicket(ticketAmount *uint8, remainingTickets uint8) {
	var userInput uint8

	fmt.Println("Enter the amount of tickets")
	fmt.Scan(&userInput)
	fmt.Println()

	if userInput > remainingTickets {
		fmt.Printf("There are only %v tickets available, please insert a valid amount.\n", remainingTickets)

		setTicket(ticketAmount, remainingTickets)
	} else {
		*ticketAmount = userInput
	}

}
