package src

import "fmt"

/*
	Type Declarations
	- Type declarations adalah kemampuan membuat ulang tipe data baru dari tipe data yang sudah ada
	- Type declarations biasanya digunakan untuk membuat alias terhadap tipe data yang sudah ada, dengan tujuan agar lebih mudah dimengerti
*/

func TypeDeclaration() {
	// Kode Program Type Declarations
	type NoKTP string
	type Married bool

	var ktpSandy NoKTP = "1111111111"
	var marriedStatus Married = true
	fmt.Println(ktpSandy)
	fmt.Println(NoKTP("2222222222"))
	fmt.Println(marriedStatus)
}
