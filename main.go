package main

import "fmt"

func main() {
	conferenceName := "Go conference"
	const conferenceTickets uint8 = 50
	var remainingTickets uint8 = 50

	var bookings []string

	fmt.Printf("Welcome to %v our booking app\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are availabe.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

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

	fmt.Println("Enter the amount of tickets")
	fmt.Scan(&userTickets)
	fmt.Println()

	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for bookin %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining\n", remainingTickets)
	fmt.Println(bookings)
}
