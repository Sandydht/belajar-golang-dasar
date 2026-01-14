package main

import "fmt"

func main() {
	// name string
	var name string

	name = "Sandy Dwi Handoko Trapsilo"
	fmt.Println(name)

	name = "Budi Setiawan"
	fmt.Println(name)

	// friendName string
	var friendName = "Budi Setiawan"
	fmt.Println(friendName)

	// age int32
	var age = 30
	fmt.Println(age)

	// country string
	country := "Indonesia"
	fmt.Println(country)

	var (
		firstName = "Sandy"
		lastName  = "Dwi"
	)
	fmt.Println("first name = ", firstName)
	fmt.Println("last name = ", lastName)
}
