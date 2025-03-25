package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	bookings := []string{}

	//fmt.Printf("ConfereceTickets are %T, Remaining Tickets are %T, Conference Name is %T \n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application \n", conferenceName) // this is a comment our conference booking application"
	fmt.Printf("We have total of %v tickets and %v are still available \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
		var userFirstName string
		var userLastName string
		var userEmail string
		var userTickets uint

		// ask user for their name
		fmt.Println("Enter your name: ")
		fmt.Scan(&userFirstName)

		fmt.Println("Enter your last Name:")
		fmt.Scan(&userLastName)

		fmt.Println("Enter your Email:")
		fmt.Scan(&userEmail)

		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		//First and Last Names Validation
		isValidName := len(userFirstName) >= 2 && len(userLastName) >= 2
		//Email Validation
		isValidEmail := strings.Contains(userEmail, "@")
		//Number of Tickets Validation
		isValidNumberOfTickets := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidNumberOfTickets {
			remainingTickets = remainingTickets - uint(userTickets)
			bookings = append(bookings, userFirstName+" "+userLastName)

			// fmt.Printf("The whole slice: %v \n", bookings)
			// fmt.Printf("The first value: %v \n", bookings[0])
			// fmt.Printf("Slice type: %T \n", bookings)
			// fmt.Printf("Slice length: %v \n", len(bookings))

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v \n", userFirstName, userLastName, userTickets, userEmail)
			fmt.Printf(("%v tickets remaining for %v \n"), remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("These are all our bookings: %v\n", firstNames)

			if remainingTickets == 0 {
				// end program
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			fmt.Printf("Your input data is invalid, Try again")
		}
	}
}
