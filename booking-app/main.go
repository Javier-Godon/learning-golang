package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const availableConferenceTickets uint = 50
	const errorMessage = "Ooops error"
	var bookings [50]string
	var sliceBookings []string
	var interactionsCounter = 0

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

	bookings[interactionsCounter] = fmt.Sprintf("%v %v", firstName, lastName)
	interactionsCounter += 1
	strings := append(sliceBookings, fmt.Sprintf("%v %v", firstName, lastName))

	fmt.Printf("Hello %v %v\n", firstName, lastName)
	remainingTickets -= userTickets
	fmt.Printf("%v you have bought %v tickets\n", firstName, userTickets)
	fmt.Printf("They are still available %v tickets for %v \n", remainingTickets, conferenceName)

	fmt.Printf("The whole array: %v\n", bookings)
	fmt.Printf("The first element: %v\n", bookings[0])
	fmt.Printf("The type: %T\n", bookings)
	fmt.Printf("The length of the array: %v\n", len(bookings))

	fmt.Printf("The whole slice: %v\n", strings)
	fmt.Printf("The first element: %v\n", strings[0])
	fmt.Printf("The type: %T\n", strings)
	fmt.Printf("The length of the slice: %v\n", len(strings))

	//https://www.youtube.com/watch?v=yyUHQIec83I
	//min 1:11:15

}
