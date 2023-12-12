package main

import "fmt"

func main() {
	var conferenceName string = "Go conference"
	const conferenceTickets uint8 = 50
	var remainingTickets uint8 = 50

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T\n", conferenceTickets, remainingTickets)
	fmt.Println()

	fmt.Printf("Welcome to %v our booking app\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are availabe.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var username string
	var userTickets int

	username = "John Doe Da Goat"
	userTickets = 2

	fmt.Println()

	fmt.Printf("User %v booked %v tickets\n", username, userTickets)
}
