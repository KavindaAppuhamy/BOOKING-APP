package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	fmt.Printf("ConfereceTickets are %T, Remaining Tickets are %T, Conference Name is %T \n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application \n", conferenceName) // this is a comment our conference booking application"
	fmt.Printf("We have total of %v tickets and %v are still available \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var userFirstName string
	var userLastName string
	var userEmail string
	var userTickets int

	// ask user for their name
	fmt.Println("Enter your name: ")
	fmt.Scan(&userFirstName)

	fmt.Println("Enter your last Name:")
	fmt.Scan(&userLastName)

	fmt.Println("Enter your Email:")
	fmt.Scan(&userEmail)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v \n", userFirstName, userLastName, userTickets, userEmail)
}
