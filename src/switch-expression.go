package src

import "fmt"

/**
Switch expression
- Selain if expression, untuk melakukan percabangan, kita juga bisa menggunakan switch expression
- Switch expression sangat sederhana dibandingkan if
- Biasanya switch expression digunakan untuk melakukan pengecekan ke kondisi dalam satu variable
*/

/**
Switch dengan Short Statement
- Sama dengan if, switch juga mendukung short statement
*/

/**
Switch Tanpa Kondisi
- Kondisi di switch expression tidak wajib
- Jika kita tidak menggunakan kondisi di switch expression, kita bisa menambahkan kondisi tersebut di setiap case nya
*/

func SwitchExpression() {
	// Kode Program Switch
	name := "Sandy"

	switch name {
	case "Sandy":
		fmt.Println("Hello Sandy")
	case "Dwi":
		fmt.Println("Hello Dwi")
	default:
		fmt.Println("Hello what is your name ?")
	}

	// Kode Program Switch dengan Short Statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	// Kode Program Switch Tanpa Kondisi
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
