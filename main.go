package main

import "fmt"

func main() {
	var conferenceName = "Go conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("Welcome to %v our booking app\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are availabe.\n", conferenceTickets,remainingTickets)
	fmt.Println("Get your tickets here to attend")

}
