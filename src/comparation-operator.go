package src

import "fmt"

/*
	Operasi Perbandingan
	- Operasi perbandingan adalah operasi untuk membandingkan dua buah data
	- Operasi perbandingan adalah operasi yang menghasilkan nilai boolean (benar atau salah)
	- Jika hasil operasinya adalah benar, maka nilainya adalah true
	- Jika hasil operasinya adalah salah, maka nilainya adalah false
*/

/*
	Operator Perbandingan
	- > -> lebih dari
	- < -> kurang dari
	- >= -> lebih dari sama dengan
	- <= -> kurang dari sama dengan
	- == -> sama dengan
	- != -> tidak sama dengan
*/

func ComparationOperator() {
	// Kode Program Operasi Perbandingan
	var name1 string = "Sandy"
	var name2 string = "Dwi"
	var result bool = name1 == name2
	fmt.Println("name1 == name2 =", result)

	var value1 = 100
	var value2 = 200
	fmt.Println("value1 > value2 =", value1 > value2)
	fmt.Println("value1 < value2 =", value1 < value2)
	fmt.Println("value1 == value2 =", value1 == value2)
	fmt.Println("value1 != value2 =", value1 != value2)
}
