package src

import "fmt"

/**
Pass by Value
- Secara default di Go-Lang semua variable itu di passing by value, bukan by reference
- Artinya, jika kita mengirim sebuah variable ke dalam function, method atau variable lain, sebenarnya yang dikirim adalah duplikasi value nya
*/

/**
Pointer
- Pointer adalah kemampuan membuat reference ke lokasi data di memory yang sama, tanpa menduplikasi data yang sudah ada
- Sederhananya, dengan kemampuan pointer, kita bisa membuat pass by reference
*/

/**
Operator &
- Untuk membuat sebuah variable dengan nilai pointer ke variable yang lain, kita bisa menggunakan operator & diikuti dengan nama variable nya
*/

/**
Operator *
- Saat kita mengubah variable pointer, maka yang berubah hanya variable tersebut
- Semua variable yang mengacu ke data yang sama tidak akan berubah
- Jika kita ingin mengubah seluruh variable yang mengacu ke data tersebut, kita bisa menggunakan operator *
*/

/**
Function New
- Sebelumnya untuk membuat pointer, kita bisa menggunakan operator &
- Go-Lang juga memiliki function new yang bisa digunakan untuk membuat pointer
- Namun function new hanya mengembalikan pointer ke data kosong, artinya tidak ada data awal
*/

/*
*
Pointer di function
- Saat kita membuat parameter di function, secara default adalah pass by value, artinya data akan di copy lalu dikirim ke function tersebut
- Oleh karena itu, jika kita mengubah data di dalam function, data yang aslinya tidak akan pernah berubah
- Hal ini membuat variable menjadi aman, karena tidak akan bisa diubah
- Namun kadang kita ingin membuat function yang bisa mengubah data asli parameter tersebut
- Untuk melakukan ini, kita juga bisa menggunakan pointer di function
- Untuk menjadikan sebuah parameter sebagai pointer, kita bisa menggunakan operator * di parameternya
*/

type Address struct {
	City, Province, Country string
}

func ChangeAddressToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func Pointer() {
	// Kode Program Pointer di Variable
	address1 := Address{
		City:     "Subang",
		Province: "Jawa Barat",
		Country:  "Indonesia",
	}
	address2 := &address1 // data di address1 di duplikat ke address2 (pass by value)
	address3 := &address1

	address2.City = "Bandung"
	*address2 = Address{
		City:     "Jakarta",
		Province: "DKI Jakarta",
		Country:  "Indonesia",
	}

	fmt.Println(address1) // address1 tidak berubah
	fmt.Println(address2)
	fmt.Println(address3)

	address4 := new(Address)
	address4.City = "Bandung"
	fmt.Println(address4)

	// Kode Program Pointer di Function
	alamat := Address{
		City:     "Subang",
		Province: "Jawa Barat",
		Country:  "",
	}
	alamatPointer := &alamat
	ChangeAddressToIndonesia(alamatPointer)
	fmt.Println(alamat)
}
