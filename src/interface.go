package src

import "fmt"

/**
Interface
- Interface adalah tipe data abstract, dia tidak memiliki implementasi langsung
- Sebuah interface berisikan definisi - definisi method
- Biasanya interface digunakan sebagai kontrak
*/

/**
Implementasi Interface
- Setiap tipe data yang sesuai dengan kontrak interface, secara otomatis dianggap sebagai interface tersebut
- Sehingga kita tidak perlu mengimplementasikan interface secara manual
- Hal ini agak berbeda dengan bahasa pemrograman lain yang ketika membuat interface, kita harus menyebutkan secara eksplisit akan menggunakan interface mana
*/

// Kode Program Interface (1)
type HasName interface {
	GetName() string
}

func SayGoodBye(hasName HasName) {
	fmt.Println("Good bye", hasName.GetName())
}

type Employee struct {
	Name string
}

func (employee Employee) GetName() string {
	return employee.Name
}

// Kode Program Interface (2)
type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func Interface() {
	var sandy Employee
	sandy.Name = "Sandy"

	var cat Animal
	cat.Name = "Milo"

	SayGoodBye(sandy)
	SayGoodBye(cat)
}
