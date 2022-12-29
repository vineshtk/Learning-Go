package helper

import "strings"

func UserInputValidation(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := remainingTickets >= userTickets && userTickets > 0

	return isValidName, isValidEmail, isValidTicketNumber
}
