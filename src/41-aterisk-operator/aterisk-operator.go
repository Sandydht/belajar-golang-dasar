package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Kendal", "Jawa Tengah", "Indonesia"}
	address2 := &address1

	address2.City = "Jakarta"

	*address2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"}

	fmt.Println(address1) // address 1 tidak berubah
	fmt.Println(address2)
}
