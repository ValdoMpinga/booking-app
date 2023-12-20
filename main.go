package main

import (
	"booking-app/helper"
	"fmt"
	"strconv"
)

var conferenceName string = "Go conference"

const conferenceTickets uint8 = 50

var remainingTickets uint8 = 50
var bookings = make([]map[string]string,0)

func main() {

	greetUsers()

	for {

		var firstName, lastName, email, userTickets = getUserInput()

		var isValidName, isValidEmail, isValidTicketNumber = helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTickets, firstName, lastName, email)

			fmt.Printf("Booking first names: %v\n", getFirstNames())

			if remainingTickets == 0 {
				fmt.Println("Our conference has no tickets left, until next year!")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("first name or last name you entered is too short")
			}
			if isValidEmail {
				fmt.Println("email address you entered doesn't contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}
		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v our booking app\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are availabe.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getUserInput() (string, string, string, uint8) {
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

	//Handles validation for the user tickets input
	setTicket(&userTickets, remainingTickets)

	return firstName, lastName, email, userTickets
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

func bookTicket(userTickets uint8, firstName, lastName, email string) {
	remainingTickets = remainingTickets - userTickets

	//create a map for a user
	var userData = make(map[string]string)

	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["numberOfTickets"] = strconv.Itoa(int(userTickets))

	bookings = append(bookings, userData)

	fmt.Printf("List of bookings %v \n\n",bookings)

	fmt.Printf("Thank you %v %v for bookin %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func getFirstNames() []string {
	firstNames := []string{}

	for _, booking := range bookings {
		firstNames = append(firstNames,booking["firstName"])

	}
	return firstNames
}
