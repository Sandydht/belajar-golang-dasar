package src

import "fmt"

/*
	Constant
	- Constant adalah variable yang nilainya tidak bisa diubah lagi setelah pertama kali diberi nilai
	- Cara pembuatan constant mirip dengan variable, yang membedakan hanya kata kunci yang digunakan adalah const, bukan var
	- Saat pembuatan constant, kita wajib langsung menginisialisasikan datanya
*/

/*
	Deklarasi Multiple Constant
	- Sama seperti variable, di Go-Lang juga kita bisa membuat constant secara multiple
*/

func Constant() {
	const (
		firstName string = "Sandy"
		lastName  string = "Dwi"
		value     int16  = 1000
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(value)
}
