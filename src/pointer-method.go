package src

import "fmt"

/**
Pointer di Method
- Walaupun method akan menempel di struct, tapi sebenarnya data struct yang diakses di dalam method adalah pass by value
- Sangat direkomendasikan menggunakan pointer di method, sehingga tidak boros memory karena harus selalu di duplikasi ketika memanggil method
*/

// Kode Program Pointer di Method
type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func PointerMethod() {
	sandy := Man{
		Name: "Sandy",
	}
	sandy.Married()
	fmt.Println(sandy.Name)
}
