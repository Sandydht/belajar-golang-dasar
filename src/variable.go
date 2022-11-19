package src

import "fmt"

/*
	Variable
	- Variable adalah tempat untuk menyimpan data
	- Variable digunakan agar kita bisa mengakses data yang sama dimanapun kita mau
	- Di Go-Lang variable hanya bisa menyimpan tipe data yang sama, jika kita ingin menyimpan data yang berbeda - beda jenis, kita harus membuat beberapa variable
	- Untuk membuat variable, kita bisa menggunakan kata kunci var, lalu diikuti dengan nama variable dan tipe datanya
*/

/*
	Tipe Data Variable
	- Saat kita membuat variable, maka kita wajib menyebutkan tipe data variable tersebut
	- Namun jika kita langsung menginisialisasikan data pada variable nya, maka kita tidak wajib menyebutkan tipe data variable nya
*/

/*
	Kata Kunci Var
	- Di Go-Lang, kata kunci var saat membuat variable tidak lah wajib
	- Asalhkan saat membuat variable kita langsung menginisialisasi datanya
	- Agar tidak perlu menggunakan kata kunci var, kita perlu menggunakan kata kunci := saat menginisialisasikan data pada variable tersebut
*/

/*
	Deklarasi Multiple Variable
	- Di Go-Lang kita bisa membuat variable secara multiple
	- Code yang dibuat akan lebih bagus dan mudah dibaca
*/

func Variable() {
	// name string
	var name string

	name = "Sandy Dwi"
	fmt.Println(name)

	name = "Sandy Handoko"
	fmt.Println(name)

	// friendName string
	var friendName = "Budi"
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
	fmt.Println("first name =", firstName)
	fmt.Println("last name =", lastName)
}
