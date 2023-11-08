package main

import (
	"fmt"
	"strings"
	"time"
)

const conferenceTickets = 50
var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberofTickets uint
}

func main() {

	greetUser()

	for {

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidemail, isValedTicketNumber := ValidateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidemail && isValedTicketNumber {

			bookTicket(userTickets, firstName, lastName, email)
			go sendTicket(userTickets, firstName, lastName, email)

			FirstNames := getFirstNames()
			fmt.Printf("The first names of booking : %v\n", FirstNames)

			if remainingTickets == 0 {
				fmt.Println("Our Conference is booked out Please join next year")
				break
			}

		} else {
			if !isValidName {
				fmt.Println("Your Name is invalid")
			}

			if !isValidemail {
				fmt.Println("Your email is not including @ sign")
			}

			if !isValedTicketNumber {
				fmt.Println("Tickets number you entered is NOT valid")
			}
		}

	}

}

func greetUser() {
	fmt.Printf("Welcome to %v booking aplication\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v is avaliable now\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your ticket here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		//firstNames = append(firstNames, booking["firstName"])

		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func ValidateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidemail := strings.Contains(email, "@")
	isValedTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidemail, isValedTicketNumber

}

func getUserInput() (string, string, string, uint) {

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your First name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your Last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your E-mail address: ")
	fmt.Scan(&email)

	fmt.Println("Enter your ticket number: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

///create a mpa for user

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	//var userData = make(map[string]string)
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberofTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of booking is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. We will send you email at %v for confirmation.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v.\n", remainingTickets, conferenceName)

}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(20 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("###############")
	fmt.Printf("Sending ticket: \n %v to email address %v\n", ticket, email)
	fmt.Println("###############")
}
