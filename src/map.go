package src

import "fmt"

/**
Tipe Data Map
- Pada array atau slice, untuk mengakses data, kita menggunakan index number dimulai dari 0
- Map adalah tipe data lain yang berisikan kumpulan data yang sama, namun kita bisa menentukan jenis tipe data index yang akan kita gunakan
- Sederhananya, map adalah tipe data kumpulan key-value (kata kunci - nilai), dimana kata kuncinya bersifat unik (tidak boleh sama)
- Berbeda dengan array dan slice, jumlah data yang kita masukan ke dalam map boleh sebanyak banyaknya, asalkan kata kunci nya berbeda, jika kita gunakan kata kunci sama, maka secara otomatis data sebelumnya akan diganti dengan data baru
*/

/**
Function Map
- len(map) -> untuk mendapatkan jumlah data di map
- map[key] -> mengambil data di map dengan key
- map[key] = value -> mengubah data di map dengan key
- make(map[TypeKey]TypeValue) -> membuat map baru
- delete(map, key) -> menghapus data di map dengan key
*/

func Map() {
	// Kode Program Map (1)
	person := map[string]string{
		"name":    "Sandy",
		"address": "Kendal",
	}

	person["title"] = "Programmer"

	fmt.Println("person =", person)
	fmt.Println("name =", person["name"])
	fmt.Println("address =", person["address"])

	// Kode Program Map (2)
	var book map[string]string = make(map[string]string)

	book["title"] = "Belajar Golang Dasar"
	book["author"] = "Sandy"
	book["ups"] = "ups"

	fmt.Println("book =", book)
	delete(book, "ups")
	fmt.Println("book =", book)
	fmt.Println("len(book) =", len(book))
}
