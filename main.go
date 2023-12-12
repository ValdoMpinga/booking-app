package main

import "fmt"

func main() {
	var conferenceName = "Go conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Println("Welcome to our", conferenceName, "booking app")
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are availabe!")
	fmt.Println("Get your tickets here to attend")

}
