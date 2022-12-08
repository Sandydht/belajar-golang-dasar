package src

import (
	"belajar-golang-dasar/src/helper"
	"fmt"
)

/**
Access Modifier
- Di bahasa pemrograman lain, biasanya ada kata kunci yang bisa digunakan untuk menentukan access modifier terhadap suatu function atau variable
- Di Go-Lang untuk menentukan access modifier, cukup dengan nama function atau variable
- Jika nama nya diawali dengan huruf besar, maka artinya bisa diakses dari package lain, jika dimulai dengan huruf kecil, artinya tidak bisa diakses dari package lain
*/

func AccessModifier() {
	fmt.Println(helper.Application)
	fmt.Println(helper.SayHello("Sandy"))
}
