package src

import "fmt"

/**
Operasi Boolean
- && -> dan
- || -> atau
- ! -> kebalikan / negasi
*/

/**
Operasi && (dan)
- true && true = true
- true && false = false
- false && true = false
- false && false = false
*/

/**
Operasi || (atau)
- true || true = true
- true || false = true
- false || true = true
- false || false = false
*/

/**
Operasi ! (negasi)
- ! true = false
- ! false = true
*/

func BooleanOperator() {
	// Kode Program Operasi Boolean
	var ujian = 85
	var absensi = 80

	var lulusUjian = ujian > 80
	var lulusAbsensi = absensi > 80

	var lulus = lulusUjian && lulusAbsensi

	fmt.Println("lulus ujian =", lulusUjian)
	fmt.Println("lulus absensi =", lulusAbsensi)
	fmt.Println("lulus =", lulus)
}
