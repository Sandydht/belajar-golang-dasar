package src

import "fmt"

/**
Function as Parameter
- Function tidak hanya bisa kita simpan di dalam variable sebagai value
- Namun juga bisa kita gunakan sebagai parameter untuk function lain
*/

/**
Function Type Declarations
- Kadang jika function terlalu panjang, agak ribet untuk menuliskannya di dalam parameter
- Type declarations juga bisa digunakan untuk membuat alias function, sehingga akan mempermudah kita menggunakan function sebagai parameter
*/

// Kode Program Function Type Declarations
type Filter func(string) string

// Kode Program Function as Parameter
func SayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func SpamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func FunctionAsParameter() {
	SayHelloWithFilter("Hello world", SpamFilter)
	SayHelloWithFilter("Anjing", SpamFilter)
}
