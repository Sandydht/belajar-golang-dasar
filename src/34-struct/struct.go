package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var sandy Customer
	sandy.Name = "Sandy Dwi Handoko Trapsilo"
	sandy.Address = "Indonesia"
	sandy.Age = 17
	fmt.Println(sandy)

	joko := Customer{
		Name:    "Joko",
		Address: "Indonesia",
		Age:     20,
	}
	fmt.Println(joko)

	budi := Customer{"Budi", "Jakarta", 35}
	fmt.Println(budi)
}
