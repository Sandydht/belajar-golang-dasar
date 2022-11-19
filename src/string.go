package src

import "fmt"

/*
	Tipe Data String
	- String adalah kumpulan karakter
	- Jumlah karakter di dalam string bisa nol sampai tak terhingga
	- Tipe data string di Go-Lang direpresentasikan dengan kata kunci string
	- Nilai data string di Go-Lang selalu diawali dan diakhiri dengan (petik dua ")
*/

/*
	Function untuk String
	- len("string") -> menghitung jumlah karakter di string
	- "string"[number] -> mengambil karakter pada posisi yang ditentukan
*/

func String() {
	// Contoh kode program tipe data string
	fmt.Println("Tipe Data String")
	fmt.Println("Sandy")
	fmt.Println("Dwi")
	fmt.Println("Handoko")
	fmt.Println("Trapsilo")

	// Contoh kode program function untuk string
	fmt.Println("\nFunction untuk String")
	fmt.Println(`len("Sandy") =`, len("Sandy"))
	fmt.Println(`"Sandy"[0] =`, "Sandy"[0]) // Megembalikan byte dari "S" yaitu 83
}
