package main

import "fmt"

func getCompleteName() (firstName string, middleName string, lastName string) {
	firstName = "Sandy"
	middleName = "Dwi"
	lastName = "Handoko"

	return firstName, middleName, lastName
}

func main() {
	firstName, middleName, lastName := getCompleteName()
	fmt.Println(firstName, middleName, lastName)
}
