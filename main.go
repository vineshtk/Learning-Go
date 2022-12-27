// packaging the whole code in to main
package main

// importing fmt package
import "fmt"

// creating the main function
func main() {

	var conferenceName = "Go conference"
	const conferenceTickets = 50

	//simpler variable declatation
	remainingTickets := 50

	//using Ptintf
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("Total %v slots and %v slots are still available to book.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here!")

	var firstName string
	var lastName string
	var email string
	var userTickets int

	userTickets = 3

	//user can input details here this asks for user input
	fmt.Println("Enter your first name: ")
	// here '&firstName' is pointer which points toward the address of the variable 'firstName'
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Printf("Thank you %v for showing your interest \n", firstName)
	fmt.Printf("You bought %v tickets for the event \n", userTickets)

}
