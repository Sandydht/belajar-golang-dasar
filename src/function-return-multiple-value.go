package src

import "fmt"

/*
	Returning Multiple Values
	- Function tidak hanya dapat mengembalikan satu value, tapi juga bisa multiple value
	- Untuk memberitahu jika function mengembalikan multiple value, kita harus menulis semua tipe data return value nya di function tersebut
*/

/*
	Menghiraukan Return Value
	- Multiple return value wajib ditangkap semua value nya
	- Jika kita ingin menghiraukan return value tersebut, kita bisa menggunakan tanda _ (garis bawah)
*/

// Kode Program Returning Multiple Values
func GetFullName() (string, string, string) {
	return "Sandy", "Dwi", "Handoko"
}

func FunctionReturnMultipleValue() {
	firstName, middleName, _ := GetFullName()
	fmt.Println(firstName, middleName)
}
