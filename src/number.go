package src

import "fmt"

/*
	Tipe Data Number
	Ada dua jenis tipe data Number, yaitu:
	- Integer
	- Floating Point
*/

/*
	Tipe Data Integer (1)
	- int8 -> nilai minimum = -128, nilai maksimum = 127
	- int16 -> nilai minimum = -32768, nilai maksimum = 32767
	- int32 -> nilai minimum = -2147483648, nilai maksimum = 2147483647
	- int64 -> nilai minimum = -9223372036854775808, nilai maksimum = 9223372036854775807
*/

/*
	Tipe Data Integer (2)
	- uint8 -> nilai minimum = 0, nilai maksimum = 255
	- uint16 -> nilai minimum = 0, nilai maksimum = 65535
	- uint32 -> nilai minimum = 0, nilai maksimum = 4294967295
	- uint64 -> nilai minimum = 0, nilai maksimum = 18446744073709551615
*/

/*
	Tipe Data Floating Point
	- float32 -> nilai minimum = 1.18 * 10 pangkat -32, nilai maksimum = 3.4 * 10 pangkat 38
	- float64 -> nilai minimum = 2.23 * 10 pangkat -308, nilai maksimum = 1.80 * 10 pangkat 308
	- complex64 -> complex numbers with float32 real and imaginary parts
	- complex128 -> complex numbers with float64 real and imaginary parts
*/

/*
	Alias
	- byte -> uint8
	- rune -> int32
	- int -> minimal int32
	- uint -> minimal uint32
*/

/*
	Catatan
	- Semakin besar kapasitasnya semakin besar penggunaannya di memori
	- Penggunaan tipe data di Go-Lang itu case sensitif
*/

func Number() {
	// Contoh kode program integer
	fmt.Println("Integer")
	fmt.Println("Satu =", 1)
	fmt.Println("Dua =", 2)

	// Contoh kode program floating point
	fmt.Println("\nFloating Point")
	fmt.Println("Tiga Koma Lima =", 3.5)
}
