package helper

import "strings"

//while running multiple files, execute the command go run . : which will run all the files in the same package

//Three types of scoping in GO: local, package, global

func ValidateUserInput(firstName string, lastName string, email string, userTickets int, remainingTickets uint, city string) (bool, bool, bool, bool) { //returning multiple values
	//applying conditions in firstname and lastname

	isValidName := len(firstName) >= 2 && len(lastName) >= 2

	//applying conditions in email
	isValidEmail := strings.Contains(email, "@")

	ValidTickets := userTickets > 0 && userTickets <= int(remainingTickets)

	isValidCity := city == "Singapore" || city == "london"

	//we can return multiple values in functions in Golang
	return isValidName, isValidEmail, ValidTickets, isValidCity
}

//we need to export the functions so that it can be used in other packages
//to do so, have the first letter of the function as capital letter, the same first letter capitalization is applicable for exporting variables too
