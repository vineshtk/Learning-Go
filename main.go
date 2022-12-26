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

	var userName string //user can input details here
	var userTickets int

	userName = "Tommy"
	userTickets = 3

	fmt.Printf("Thank you %v for showing your interest \n", userName)
	fmt.Printf("You bought %v tickets for the event \n", userTickets)

	//getting the variable type
	fmt.Printf("The type of variable conferenceName is %T \n", conferenceName)

}
