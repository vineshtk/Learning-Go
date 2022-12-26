package main

import "fmt"

func main() {

	var conferenceName = "Go conference"
	const conferenceTickets = 50
	var remainingTickets = 50
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("Total %v slots and %v slots are still available to book.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here!")

}
