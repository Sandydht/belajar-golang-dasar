package src

import "fmt"

/**
Tipe Data Slice
- Tipe data slice adalah potongan dari data array
- Slice mirip dengan array, yang membedakan adalah ukuran slice bisa berubah
- Slice dan array selalu terkoneksi, dimana slice adalah data yang mengakses sebagian atau seluruh data di array
*/

/**
Detail Tipe Data Slice
Tipe data slice memiliki 3 data, yaitu:
- Pointer -> petunjuk data pertama di array pada slice
- Length -> panjang dari slice
- Capacity -> kapasitas dari slice, dimana length tidak boleh lebih dari capacity
*/

/**
Membuat Slice Dari Array
- array[low:high] -> membuat slice dari array dimulai index low sampai index sebelum high
- array[low:] -> membuat slice dari array dimulai index low sampai index akhir di array
- array[:high] -> membuat slice dari array dimulai index 0 sampai index sebelum high
- array[:] -> membuat slice dari array dimulai index 0 sampai index akhir di array
*/

/**
Function Slice
- len(slice) -> untuk mendapatkan panjang
- cap(slice) -> untuk mendapatkan kapasitas
- append(slice, data) -> membuat slice baru dengan menambah data ke posisi terakhir slice, jika kapasitas sudah penuh, maka akan membuat array baru
- make([]TypeData, length, capacity) -> membuat slice baru
- copy(destination, source) -> menyalin slice dari source ke destination
*/

/*
	Catatan
	- Saat membuat array, kita harus berhati - hati, jika salah, maka yang kita buat bukanlah array, melainkan slice
*/

func Slice() {
	// Kode Program Slice (1)
	var months = [...]string{
		"Januari",
		"Febuari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}
	var monthSlice1 = months[4:7] // pointer = 4, length = 3, capacity = 8
	var monthSlice2 = months[6:9] // pointer = 6, length = 3, capacity = 6

	fmt.Println("months =", months)
	fmt.Println("monthSlice1 =", monthSlice1)
	fmt.Println("monthSlice2 =", monthSlice2)

	// Kode Program Slice (2)
	var days = [...]string{
		"Senin",
		"Selasa",
		"Rabu",
		"Kamis",
		"Jumat",
		"Sabtu",
		"Minggu",
	}
	var daySlice1 = days[5:]
	daySlice1[0] = "Sabtu baru"
	daySlice1[1] = "Minggu baru"
	fmt.Println("days =", days)
	fmt.Println("daySlice1 =", daySlice1)

	daySlice2 := append(daySlice1, "Libur baru")
	daySlice2[0] = "Ups"
	fmt.Println("daySlice2 =", daySlice2)
	fmt.Println("days =", days)
	fmt.Println("daySlice1 =", daySlice1)

	// Kode Program Slice (3)
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Sandy"
	newSlice[1] = "Dwi"
	fmt.Println("new slice =", newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	// Kode Program Slice (4)
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println("copy slice =", copySlice)

	// Kode Program Slice (5)
	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	// iniArray[5] = 6 // Ukuran array tidak bisa diubah
	// iniSlice[5] = 7 // Ukuran slice bisa diubah

	fmt.Println("ini array =", iniArray)
	fmt.Println("ini slice =", iniSlice)
}
