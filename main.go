package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

var title string = "DingDong Concert"

const totalTickets int = 100

var remainingTickets uint = 100

type UserData struct {
	firstName string
	lastName  string
	email     string
	tickets   uint
}

var bookings = make([]UserData, 0)
var wg = sync.WaitGroup{}

func greetUsers() {
	fmt.Printf("Hello Welcome to %v, We have %v tickets available.\nMake sure to book yours before its too late \n", title, remainingTickets)
	fmt.Println("##############################")
}

func getDetails() (string, string, string, uint) {
	var (
		firstName string
		lastName  string
		email     string
		tickets   uint
	)

	fmt.Println("Please enter your First Name : ")
	fmt.Scanf("%s", &firstName)

	fmt.Println("Please enter your Last Name : ")
	fmt.Scanf("%s", &lastName)

	fmt.Println("Please enter your Email : ")
	fmt.Scanf("%s", &email)

	fmt.Println("Please enter the number of Tickets : ")
	fmt.Scanln(&tickets)

	return firstName, lastName, email, tickets

}

func ValidDetails(firstName, lastName, email string, tickets uint) (bool, bool, bool) {
	isValidName := len(firstName) > 0 && len(lastName) > 0 && strings.IndexFunc(firstName, isNotLetter) == -1 && strings.IndexFunc(lastName, isNotLetter) == -1

	isValidEmail := strings.Contains(email, "@") && strings.Contains(email, ".")

	isValidTickets := tickets > 0 && tickets <= remainingTickets

	return isValidName, isValidEmail, isValidTickets
}

func isNotLetter(r rune) bool {
	return r < 'A' || (r > 'Z' && r < 'a') || r > 'z'
}

func bookTicket(firstName, lastName, email string, tickets uint) {
	remainingTickets = remainingTickets - tickets

	userData := UserData{
		firstName: firstName,
		lastName:  lastName,
		email:     email,
		tickets:   tickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, tickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, title)
}

func sentTickets(firstName, lastName, email string, tickets uint) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", tickets, firstName, lastName)
	fmt.Println("#################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("#################")

	wg.Done()

}

func main() {
	greetUsers()

	for {
		if remainingTickets == 0 {
			fmt.Println("Sorry, all tickets are sold out!")
			break
		}
		firstName, lastName, email, tickets := getDetails()
		isValidName, isValidEmail, isValidTickets := ValidDetails(firstName, lastName, email, tickets)

		if isValidName && isValidEmail && isValidTickets {
			bookTicket(firstName, lastName, email, tickets)

			wg.Add(1)
			go sentTickets(firstName, lastName, email, tickets)

			fmt.Printf("Thank you, %s %s! Your booking for %d tickets with the email %s is successful.\n", firstName, lastName, tickets, email)
		} else {
			fmt.Println("There was an error with your input:")
			if !isValidName {
				fmt.Println("- Name is invalid. Make sure it’s not empty and contains only letters.")
			}
			if !isValidEmail {
				fmt.Println("- Email is invalid. Make sure it contains '@' and a valid domain.")
			}
			if !isValidTickets {
				fmt.Println("- Number of tickets must be greater than 0.")
			}
		}

	}
	wg.Wait()

}
