package src

import "fmt"

/*
	Function Parameter
	- Saat membuat function, kadang - kadang kita membutuhkan data dari luar, atau kita sebut sebagai parameter
	- Kita bisa menambahkan parameter di function lebih dari satu
	- Parameter tidaklah wajib, jadi kita bisa membuat function tanpa parameter seperti sebelumnya yang sudah kita buat
	- Namun jika kita menambahkan parameter di function, maka ketika memanggil function tersebut, kita wajib memasukan data ke parameternya
*/

/*
	Catatan
	- Nama parameter bersifat unik
*/

// Kode Program Function Parameter
func SayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func FunctionParameter() {
	SayHelloTo("Sandy", "Dwi")
	SayHelloTo("Budi", "Domisol")
}
