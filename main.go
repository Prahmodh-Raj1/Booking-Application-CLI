//always initialize a module before using Go : go mod init <module name>
//all the code must belong to a package
//the first statement in a go program is "package main"

package main

//import other packages if we wanna use files in that package
//declaring the entry point for the application, as to where to start the program
import (
	"booking-app/helper"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
) //stands for format

//for importing user defined packages, use the path name too like booking-app/helper

// declaring the package level variables (global variables)
var conferenceName = "Go conference" //declaring the variable name with the var keyword
// fmt.Println("Name of the conference is ", conferenceName) //println is used to print
const conferenceTickets = 50   //const keyword to declare the variable whose value does not change with time
var remainingTickets uint = 50 //used to track the remaining tickets
var bookingslice []string      //declaring a slice

var wg = sync.WaitGroup{} //waits for the launched goroutine to finish
//package "sync" provides basic synchronization functionality

func main() { //similar to int main in c++
	//Print is a function that comes from a package called fmt

	greetUsers(conferenceName, conferenceTickets)
	fmt.Println("Welcome to the ", conferenceName, " booking application")
	fmt.Println("We have ", conferenceTickets, " tickets and the remaining tickets are ", remainingTickets)
	fmt.Println("Grab your tickets to attend") //println is used to print newline character along

	//Printing  formatted data; fmt.printf("Some text with a variable %s",myVariableName)
	//it takes a template string that takes the text that needs to be formatted
	fmt.Println("\nYou are competing for %v tickets", remainingTickets) //%v is used as placeholders

	//Printing the reference to a variable
	fmt.Println("Address of variable is ", &conferenceName)

	//Variables and datatypes

	var firstName string
	var lastName string
	var email string
	var userTickets int
	//we ask for username
	fmt.Printf("Enter the username: ")
	fmt.Scan(&firstName) //used to take input. The parameter is the pointer reference to the variable that we are taking input for

	fmt.Printf("Enter the last name: ")
	fmt.Scan(&lastName)
	fmt.Printf("Enter the email: ")
	fmt.Scan(&email)

	fmt.Printf("Enter the user tickets: ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - uint(userTickets) //applying type casting
	fmt.Printf("User %v %v has booked %v tickets", firstName, lastName, userTickets)
	fmt.Printf("\nFirstname has %T type", firstName) //%T is a placeholder that stores the type of the variable it stores

	fmt.Println("\nRemaining tickets is ", remainingTickets)
	//Datatypes in GO: string, integers,booleans, maps, arrays

	//Arrays and slices

	//var bookernames = [20]string{"Jason", "percy", "Frank"} //declaring the size of the array within the square brackets
	//declaring  an array of size 20 with 3 elements already in it

	var userNames [25]string //declaring an array of size 25

	userNames[0] = firstName + " " + lastName

	fmt.Println("The whole array is ", userNames)

	fmt.Println("The first username is ", userNames[0])

	fmt.Println("The length of the array is ", len(userNames))

	//when we don't know the size of the array, we can use a slice (a dynamic-sized array)

	//empty slice assignment var bookingslice []string{}
	//appending elements to the slice
	bookingslice = append(bookingslice, "Annabeth Chase") //appending elements to the slice

	fmt.Println("The whole slice is ", bookingslice)

	fmt.Println("The first element of bookingslice is ", bookingslice[0])

	fmt.Println("Size of the slice is ", len(bookingslice))

	//LOOPS:
	//we have only the for loop in the go language

	for remainingTickets > 0 && len(bookingslice) < 50 { //this is how we set conditions and stop the execution of the loop
		//for true {} can be considered similar to while(true) in other languages

		firstName, lastName, email, userTickets, city := getUserInput()
		isValidName, isValidEmail, ValidTickets, isValidCity := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets, city)
		if isValidName && isValidEmail && ValidTickets && isValidCity {
			bookingslice = append(bookingslice, firstName+" "+lastName)
			remainingTickets = remainingTickets - uint(userTickets) //applying type casting

			fmt.Printf("User %v %v has booked %v tickets", firstName, lastName, userTickets)
			fmt.Println("\nRemaining tickets is ", remainingTickets)

			//right before we spawn a new thread, we call the Add function of WaitGroup
			wg.Add(1) //Add sets the number of go routines to wait for  (increases the counter by the provided number)

			go sendTickets(uint(userTickets), firstName, email) //achieving concurrency by using the go keyword allowing to work as multiple threads

			firstNames := printFirstNames(bookingslice)
			fmt.Printf("The first names of bookings are: %v\n", firstNames)
			//Conditionals

			if remainingTickets == 0 {
				//code to be executed
				fmt.Println("Conference is booked out mate")
				break //breaks out of the loop
			}
		} else if userTickets == int(remainingTickets) {
			//do smtg here
		} else {
			if !isValidCity {
				fmt.Println("Your city is not valid")
			}
			if !isValidEmail {
				fmt.Println("Your email does not contain @")
			}
			if !isValidName {
				fmt.Println("Invalid name")
			}
			if !ValidTickets {
				fmt.Println("Tickets booking out of bounds")
			}

			continue
		}
		wg.Wait() //waits for all the threads to complete , i.e blocks until the waitGroup counter is 0
	}

	confcity := "Mountain View"
	switch confcity {
	case "New York":
		//execute code for booking in NYC
		break
	case "Mountain View", "Bay Area":
		//execute code for booking in Mountain View and Bay area
		break
	case "Berlin", "Munich": //for multiple cases that have same code block
		//execute code for booking in Berlin
		break
	default:
		fmt.Println("Invalid city")
		break
	}

	//For each loop
}

func greetUsers(conferenceName string, confTickets int) {
	fmt.Println("Welcome  to the conference of Golang developers and ", conferenceName, " Available tickets: ", confTickets)
}

/*
func validateUserInput(firstName string, lastName string, email string, userTickets int, remainingTickets uint, city string) (bool, bool, bool, bool) { //returning multiple values

		//applying conditions in firstname and lastname

		isValidName := len(firstName) >= 2 && len(lastName) >= 2

		//applying conditions in email
		isValidEmail := strings.Contains(email, "@")

		ValidTickets := userTickets > 0 && userTickets <= int(remainingTickets)

		isValidCity := city == "Singapore" || city == "london"

		//we can return multiple values in functions in Golang
		return isValidName, isValidEmail, ValidTickets, isValidCity
	}
*/
func printFirstNames(bookingslice []string) []string { //we write the return type outside the brackets
	////for arrays, range gives the index and the value for each element  in the for each loop
	//_ is the blank identifier: to ignore a variable u don't want to use

	firstnames := []string{}
	//for each loop
	for _, first := range bookingslice { //this loop ends when iterated over all elements of the bookingslice list
		var names = strings.Fields(first) //strings.Fields(): splits the string with white space as separator
		var firstnm = names[0]
		firstnames = append(firstnames, firstnm)
	} //range iterates over elements for different data structures
	return firstnames
}

func getUserInput() (string, string, string, int, string) {
	var firstName string
	var lastName string
	var email string
	var userTickets int
	var city string
	//we ask for username
	fmt.Printf("Enter the username: ")
	fmt.Scan(&firstName) //used to take input. The parameter is the pointer reference to the variable that we are taking input for

	fmt.Printf("Enter the last name: ")
	fmt.Scan(&lastName)

	fmt.Printf("Enter the email: ")
	fmt.Scan(&email)

	fmt.Printf("Enter the user tickets: ")
	fmt.Scan(&userTickets)

	fmt.Printf("Enter the city of ur choice")
	fmt.Scan(&city)

	return firstName, lastName, email, userTickets, city

}

func useMaps() {

	//maps have unique keys to values
	//You can retrieve the value by using it's key letter
	//all keys have the same data type and all values have the same data type

	//create a map for a user
	var userMap = make(map[string]string) //used to create a map. make keyword is used to create empty map

	//adding values to maps
	userMap["firstName"] = "Jason"
	userMap["lastName"] = "Grace"
	userMap["email"] = "jgrace@gmail.com"
	var userTickets uint = 20
	//applying type casting from uint to string using FormatUint
	userMap["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10) //used to type cast from uint to string

	var newBookings = make([]map[string]string, 0) //creating a list/slice of type map
	newBookings = append(newBookings, userMap)

	//printing the newbookings
	fmt.Printf("newbookings using map: ", newBookings)

}

func getFirstnameFromMap(bookings []map[string]string) []string {
	firstnames := []string{}
	//for each loop
	for _, booking := range bookings { //this loop ends when iterated over all elements of the bookingslice list

		firstnames = append(firstnames, booking["firstName"])
	} //range iterates over
	return firstnames
}

// STRUCTURES in Golang
type userData struct { //creating a struct
	firstName            string
	lastName             string
	email                string
	numberOfTickets      int
	isOptedForNewsletter bool
}

func usingStruct() {
	var newBookings = make([]userData, 0)
	var userDetails = userData{
		firstName:            "Percy",
		lastName:             "Jackson",
		email:                "perseusjackson@gmail.com",
		numberOfTickets:      10,
		isOptedForNewsletter: true,
	}
	newBookings = append(newBookings, userDetails)
}

func getFirstnameFromStruct(bookings []userData) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames

}

func sendTickets(userTickets uint, firstname string, email string) {
	var Ticket = fmt.Sprintf("%v tickets for %v\n", userTickets, firstname) //used to generate a string variable : Sprintf
	time.Sleep(50 * time.Second)                                            //sleep for 10 seconds
	fmt.Println("######################")
	fmt.Printf("Sending ticket %v to email address: %v\n", Ticket, email)
	fmt.Println("######################")

	wg.Done() //signals to the wait group that this particular thread is done executing
}

//The main thread doesn't wait for the other threads to complete/even start executing once the main thread has finished executing

//In order to make the other threads to execute, we need to tell the main thread to wait for the other threads to complete

//Go uses what is called a "Green thread" (called go routines), which is an abstraction of the actual thread

//They are managed by the go runtime, we are only interacting with these high level go routines
//cheaper and lightweight
//we can run thousands of go routines without much impact on performance

//go allows channels for go routines to communicate with one another
