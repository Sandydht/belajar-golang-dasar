package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func changeAddressToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	address1 := Address{
		City:     "Subang",
		Province: "Jawa Barat",
		Country:  "Indonesia",
	}
	address2 := &address1 // data di address1 di duplikat ke address2 (pass by value)
	address3 := &address1

	address2.City = "Bandung"
	*address2 = Address{
		City:     "Jakarta",
		Province: "DKI Jakarta",
		Country:  "Indonesia",
	}

	fmt.Println(address1) // address1 tidak berubah
	fmt.Println(address2)
	fmt.Println(address3)

	address4 := new(Address)
	address4.City = "Bandung"
	fmt.Println(address4)

	alamat := Address{
		City:     "Subang",
		Province: "Jawa Barat",
		Country:  "",
	}
	alamatPointer := &alamat
	changeAddressToIndonesia(alamatPointer)
	fmt.Println(alamat)
}
