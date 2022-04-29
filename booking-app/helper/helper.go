package helper

import (
	"strings"
)

func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remaingingTickets uint) (bool, bool, bool) {

	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmial := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remaingingTickets

	return isValidName, isValidEmial, isValidTicketNumber
}
