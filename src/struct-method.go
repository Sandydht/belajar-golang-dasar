package src

import "fmt"

/**
Struct Method
- Struct adalah tipe data seperti tipe data lainnya, dia bisa digunakan sebagai parameter untuk function
- Namun jika kita ingin menambahkan method ke dalam struct, sehingga seakan - akan sebuah struct memiliki function
- Method adalah function
*/

// Kode Program Struct Method
type Person struct {
	Name, Address string
	Age           int
}

func (customer Customer) SayHello() {
	fmt.Println("Hello, My name is", customer.Name)
}

func StructMethod() {
	rully := Customer{Name: "Rully"}
	rully.SayHello()
}
