// packaging the whole code in to main
package main

// importing fmt package
import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

var conferenceName string = "Go conference"

const conferenceTickets = 50

var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

// creating the main function
func main() {

	for {

		// function call for greeting users
		greetUser()
		firstName, lastName, email, userTickets := getUserInputs()

		// user input validation
		isValidName, isValidEmail, isValidTicketNumber := helper.UserInputValidation(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTickets, firstName, lastName, email)

			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, email)

			// function call for printing firstnames of users
			firstNames := getFirstNames()
			fmt.Printf("The first names are: %v \n", firstNames)

			noTicketsRemaining := remainingTickets == 0
			if noTicketsRemaining {
				println("Conference is booked out.Come back next year.")
				break
			}
		} else {

			if !isValidName {
				fmt.Println("First name and Last name must contain minimum two charecters! ")
			}
			if !isValidEmail {
				fmt.Println("Email entered is in wrong format!")
			}
			if !isValidTicketNumber {
				fmt.Println("Invalid ticket number!")
			}
		}
	}

	wg.Wait()
}

func greetUser() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("Total %v slots and %v slots are still available to book.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here!")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInputs() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

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

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	// create map for user data

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("list of bookings %v\n \n", bookings)
	fmt.Printf("Thank you %v %v for showing your interest. You booked %v tickets,you will receive conformation email at %v shortly \n", firstName, lastName, userTickets, email)

}

func sendTicket(userTicket uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v Tickets for %v %v", userTicket, firstName, lastName)
	fmt.Println("########################")
	fmt.Printf("Sending ticket:\n%v \nto email address %v\n", ticket, email)
	fmt.Println("########################")
	wg.Done()
}
