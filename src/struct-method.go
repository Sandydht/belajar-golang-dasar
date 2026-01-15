package src

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

type Person struct {
	Name, Address string
	Age           int
}

func (customer Customer) SayHello() {
	fmt.Println("Hello, My name is", customer.Name)
}

func StructMethod() {
	rully := Customer{Name: "Rully"}
	rully.SayHello()
}
