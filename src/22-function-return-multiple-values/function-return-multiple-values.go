package main

import "fmt"

func getFullName() (string, string, string) {
	return "Sandy", "Dwi", "Handoko"
}

func main() {
	firstName, middleName, lastName := getFullName()
	fmt.Println(firstName, middleName, lastName)
}
