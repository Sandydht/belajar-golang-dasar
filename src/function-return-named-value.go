package src

import "fmt"

/**
Named Return Values
- Biasanya saat kita memberi tahu bahwa sebuah function mengembalikan value, maka kita hanya mendeklarasikan tipe data return value di function
- Namun kita juga bisa membuat variable secara langsung di tipe data return function nya
*/

// Kode Program Named Return Values
func GetCompleteName() (firstName string, middleName string, lastName string) {
	firstName = "Sandy"
	middleName = "Dwi"
	lastName = "Handoko"

	return
}

func FunctionReturnNamedValue() {
	firstName, middleName, lastName := GetCompleteName()
	fmt.Println(firstName, middleName, lastName)
}
