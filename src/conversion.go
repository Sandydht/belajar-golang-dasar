package src

import "fmt"

/**
Konversi Tipe Data
- Di Go-Lang kadang kita butuh melakukan konversi tipe data dari satu tipe ke tipe lain
- Misal kita ingin mengkonversi tipe data int32 ke int64, dan lain - lain
*/

/**
Catatan
- Berhati - hati dalam melakukan konversi tipe data, perhatikan kapasitas, ex: nilai32 = 32768 di konversi menjadi tipe data int8, maka akan kembali ke nilai minimum int8
*/

func Conversion() {
	// Kode Program Konversi Tipe Data (1)
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)
	var nilai8 int8 = int8(nilai32)
	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)
	fmt.Println(nilai8)

	// Kode Program Konversi Tipe Data (2)
	var name string = "Sandy Dwi"
	var e = name[0]         // Mengembalikan byte dari "S"
	var eString = string(e) // Mengkonversi dari byte "S" menjadi string "S"
	fmt.Println(name)
	fmt.Println(eString)
}
