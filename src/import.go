package src

import (
	"belajar-golang-dasar/src/helper"
	"belajar-golang-dasar/src/other"
	"fmt"
)

/**
Package
- Package adalah tempat yang bisa digunakan untuk mengorganisir kode program yang kita buat di Go-Lang
- Dengan menggunakan package, kita bisa merapikan kode program yang kita buat
- Package sendiri sebenarnya hanya direktori folder di sistem operasi kita
*/

/**
Import
- Secara standar, file Go-Lang hanya bisa mengakses file Go-Lang lainnya yang berada dalam package yang sama
- Jika kita ingin mengakses file Go-Lang yang berada diluar package, maka kita bisa menggunakan import
*/

func Import() {
	helperSayHello := helper.SayHello("Sandy")
	fmt.Println(helperSayHello)

	otherSayHello := other.SayHello("Sandy")
	fmt.Println(otherSayHello)
}
