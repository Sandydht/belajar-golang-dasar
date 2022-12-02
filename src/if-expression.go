package src

import "fmt"

/**
If Expression
- If adalah salah satu kata kunci yang digunakan untuk percabangan
- Percabangan artinya kita bisa mengeksekusi kode program tertentu ketika suatu kondisi terpenuhi
- Hampir di semua bahasa pemrograman mendukung if expression
*/

/**
Else Expression
- Blok if akan dieksekusi ketika kondisi if bernilai true
- Kadang kita ingin melakukan eksekusi program tertentu jika kondisi if bernilai false
- Hal ini bisa dilakukan menggunakan else expression
*/

/**
Else If Expression
- Kadang dalam if, kita butuh membuat beberapa modul
- Kasus seperti ini, kita bisa menggunakan else if expression
*/

/**
If dengan Short Statement
- If mendukung short statement sebelum kondisi
- Hal ini sangat cocok untuk membuat statement yang sederhana sebelum melakukan pengecekan terhadap kondisi
*/

func IfExpression() {
	// Kode Program If
	name := "Sandy"

	if name == "Sandy" {
		fmt.Println("Hello Sandy")
	} else if name == "Dwi" {
		fmt.Println("Hello Dwi")
	} else {
		fmt.Println("Hello what is your name ?")
	}

	// Kode Program If Short Statement
	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
