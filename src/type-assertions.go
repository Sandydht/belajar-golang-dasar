package src

import "fmt"

/**
Type Assertions
- Type assertions merupakan kemampuan merubah tipe data menjadi tipe data yang diinginkan
- Fitur ini sekali digunakan ketika kita bertemu dengan data interface kosong
*/

/**
Type Assertions Menggunakan Switch
- Saat salah menggunakan type assertions, maka bisa berakibat terjadi panic di aplikasi kita
- Jika panic dan tidak ter - recover, maka otomatis program kita akan mati
- Agar lebih aman, sebaiknya kita menggunakan switch expression untuk melakukan type assertions
*/

// Kode Program Type Assertions
func random() interface{} {
	return "OK"
}

func TypeAssertions() {
	result := random()
	// resultString := result.(string)
	// fmt.Println(resultString)

	// resultInt := result.(int) // panic
	// fmt.Println(resultInt)

	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknown type")
	}
}
