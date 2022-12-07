package src

import "fmt"

/**
Struct
- Struct adalah sebuah template data yang digunakan untuk menggabungkan nol atau lebih data lainnya dalam satu kesatuan
- Struct biasanya digunakan untuk merepresentasikan data dalam program aplikasi yang kita buat
- Data di struct disimpan dalam field
- Sederhananya struct adalah kumpulan dari field
*/

/**
Membuat Data Struct
- Struct adalah template data atau prototype data
- Struct tidak bisa langsung digunakan
- Namun kita bisa membuat data / object dari struct yang telah kita buat
*/

/**
Struct Literals
- Sebelumnya kita telah membuat data dari struct, namun sebenarnya ada banyak cara yang bisa kita gunakan untuk membuat data dari struct
*/

// Kode Program Struct
type Customer struct {
	Name, Address string
	Age           int
}

func Struct() {
	var sandy Customer
	sandy.Name = "Sandy Dwi Handoko Trapsilo"
	sandy.Address = "Indonesia"
	sandy.Age = 17
	fmt.Println(sandy)

	// Kode Program Struct Literals
	joko := Customer{
		Name:    "Joko",
		Address: "Indonesia",
		Age:     20,
	}
	fmt.Println(joko)

	budi := Customer{"Budi", "Jakarta", 35}
	fmt.Println(budi)
}
