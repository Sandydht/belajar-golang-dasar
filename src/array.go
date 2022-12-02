package src

import "fmt"

/**
Tipe Data Array
- Array adalah tipe data yang berisikan kumpulan data dengan tipe yang sama
- Saat membuat array, kita perlu menentukan jumlah data yang bisa ditampung oleh array tersebut
- Kapasitas array tidak bisa bertambah setelah array dibuat
*/

/**
Index di Array
- index ke 0 -> Sandy
- index ke 1 -> Dwi
- index ke 2 -> Handoko
- index ke 3 -> Trapsilo
*/

/**
Membuat Array Langsung
- Di Go-Lang kita juga bisa membuat array secara langsung saat deklarasi variable
*/

/**
Function Array
- len(array) -> untuk mendapatkan panjang array atau kapasitas array, bukan jumlah data dari array
- array[index] -> mendapatkan data di posisi index
- array[index] = value -> mengubah data di posisi index
*/

func Array() {
	// Kode Program Array (1)
	var names [4]string
	names[0] = "Sandy"
	names[1] = "Dwi"
	names[2] = "Handoko"
	names[3] = "Trapsilo"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])
	fmt.Println(names[3])
	// fmt.Println(names[4]) // invalid argument: index 4 out of bounds [0:4], karena kapasitas di array names adalah 4 yang dimulai dari index 0

	// Kode Program Array (2)
	var values = [3]int{90, 80, 95}
	fmt.Println(values)

	// Kode Program Array (3)
	fmt.Println("len(names) =", len(names))
	fmt.Println("len(values) =", len(values))
}
