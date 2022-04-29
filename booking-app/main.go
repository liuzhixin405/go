package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

//var bookings = make([]map[string]string, 0)
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {
	var conferenceName = "Go conference"
	const conferenceTickets int = 50
	var remaingingTickets uint = 50

	greetUsers(conferenceName, conferenceTickets, remaingingTickets)

	firstName, lastName, email, userTickets := getUserInput()

	//isValidCity := city != "Singapore" && city != "London"
	isValidName, isValidEmial, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remaingingTickets)
	if isValidName && isValidEmial && isValidTicketNumber {
		bookTicket(remaingingTickets, userTickets, firstName, lastName, email, conferenceName)
		//sendTicket(userTickets, firstName, lastName, email)
		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)
		var firstNames = getFirstNames()

		fmt.Printf("the first names of bookings are %v \n", firstNames)

		if remaingingTickets == 0 {
			fmt.Println("our conference is booked out . Come back next year.")
		}
	} else {
		if !isValidName {
			fmt.Println(" first name or last name you entered is too short .")
		}
		//可单独校验
		fmt.Println("yout input is invalid ,try again: ")
	}

	wg.Wait()
}

func greetUsers(conferenceName string, confTickets int, remaingingTickets uint) {
	fmt.Printf("welcome to %v booking application\n", conferenceName)
	fmt.Println("we have total of", confTickets, "tickets and", remaingingTickets, "are still avallable.")
	fmt.Println("get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {

		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var userTickets uint
	var lastName string
	var email string

	fmt.Println("enter your first name:")
	fmt.Scan(&firstName)
	fmt.Println("enter your last name:")
	fmt.Scan(&lastName)
	fmt.Println("enter your email address:")
	fmt.Scan(&email)
	fmt.Println("enter your tickets:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(remaingingTickets uint, userTickets uint, firstName string, lastName string, email string, conferenceName string) {
	remaingingTickets = remaingingTickets - userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}
	// userData["firatName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)
	if userTickets > remaingingTickets {
		fmt.Printf("we only have %v\n tickets remaing , so you can't book %v tickets \n", remaingingTickets, userTickets)
	}
	bookings = append(bookings, userData)
	fmt.Print("List of bookings is %v\n", bookings)
	fmt.Printf("thank  you %v  %v for booking  %v tickets, you will receive a confirmation email at %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaing for %v\n", remaingingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(20 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("********************")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("********************")
	wg.Done()
}
