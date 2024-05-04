package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const availableConferenceTickets uint = 50
	const errorMessage = "Ooops error"
	var remainingTickets uint = availableConferenceTickets
	fmt.Printf("Welcome to %v\n", conferenceName)
	fmt.Printf("They are still available %v tickets\n", remainingTickets)

	var firstName string
	var lastName string
	var email string
	var userTickets uint
	fmt.Println("Enter your first name: ")
	_, err := fmt.Scan(&firstName)
	if err != nil {
		fmt.Printf("%v %v", errorMessage, err)
	}

	fmt.Println("Enter your last name: ")
	_, err = fmt.Scan(&lastName)
	if err != nil {
		fmt.Printf("%v %v", errorMessage, err)
	}

	fmt.Println("Enter your email: ")
	_, err = fmt.Scan(&email)
	if err != nil {
		fmt.Printf("%v %v", errorMessage, err)
	}

	fmt.Println("Enter the number of tickets to buy: ")
	_, err = fmt.Scan(&userTickets)
	if err != nil {
		fmt.Printf("%v %v", errorMessage, err)
	}

	fmt.Printf("Hello %v %v\n", firstName, lastName)
	remainingTickets -= userTickets
	fmt.Printf("%v you have bought %v tickets\n", firstName, userTickets)
	fmt.Printf("They are still available %v tickets\n", remainingTickets)

	//min 54:03 //https://www.youtube.com/watch?v=yyUHQIec83I
}
