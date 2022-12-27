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
	var bookings [50]string // here [50] is the size of array and string is the type

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

	remainingTickets = remainingTickets - userTickets
	bookings[0] = firstName + " " + lastName

	fmt.Printf("The whole array %v\n", bookings)
	fmt.Printf("The first value %v\n", bookings[0])
	fmt.Printf("data type of array %T \n", bookings)
	fmt.Printf("the length of array %v \n", len(bookings))

	fmt.Printf("Thank you %v %v for showing your interest. You booked %v tickets,you will receive conformation email at %v shortly \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are available to book for %v \n", remainingTickets, conferenceName)
}
