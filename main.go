package main

//what functions do we have available in the fmt package?
//format package
import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	//only applies to variable, and cannot define explicit types
	remainingTickets := 50

	fmt.Printf("Welcome to %v Booking Application!\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your ticket here to attend")

	//var bookings = [50]string{}
	//slice is an abstraction of array
	var bookings []string

	fmt.Println("-----------------------------------------------")

	var firstName string
	var lastName string
	var email string
	var userTickets int

	//ask the user for their name
	fmt.Println("Enter your first Name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last Name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("User %v booked %v tickets.\n", firstName, userTickets)

	fmt.Printf("Thank you %v %v for booking %v. You will recieve a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
	fmt.Printf("These are all the bookings %v\n", bookings)

}
