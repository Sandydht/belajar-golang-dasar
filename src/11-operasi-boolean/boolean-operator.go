package main

import "fmt"

func main() {
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
