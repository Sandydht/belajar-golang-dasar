package main

import "fmt"

type HasName interface {
	getName() string
}

func sayGoodBye(hasName HasName) {
	fmt.Println("Good bye", hasName.getName())
}

type Employee struct {
	Name string
}

func (employee Employee) getName() string {
	return employee.Name
}

type Animal struct {
	Name string
}

func (animal Animal) getName() string {
	return animal.Name
}

func main() {
	var sandy Employee
	sandy.Name = "Sandy"

	var cat Animal
	cat.Name = "Milo"

	sayGoodBye(sandy)
	sayGoodBye(cat)
}
