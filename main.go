// packaging the whole code in to main
package main

// importing fmt package
import (
	"fmt"
	"strings"
)

// creating the main function
func main() {

	conferenceName := "Go conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	bookings := []string{}

	//using Ptintf
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("Total %v slots and %v slots are still available to book.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here!")

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	for {
		//user can input details here this asks for user input
		fmt.Println("Enter your first name: ")
		// here '&firstName' is pointer which points towards the address of the variable 'firstName'
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		if remainingTickets < userTickets {
			fmt.Printf("there are only %v tickets avilable to book.", remainingTickets)
			break
		}

		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %v %v for showing your interest. You booked %v tickets,you will receive conformation email at %v shortly \n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v \n", remainingTickets, conferenceName)

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("The first names are: %v \n", firstNames)

		noTicketsRemaining := remainingTickets == 0
		if noTicketsRemaining {
			println("Conference is booked out.Come back next year.")
			break
		}
	}
}
