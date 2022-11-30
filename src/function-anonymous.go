package src

import "fmt"

/*
	Anonymous Function
	- Sebelumnya setiap membuat function, kita akan selalu memberikan sebuah nama pada function tersebut
	- Namun kadang ada kalanya lebih mudah membuat function secara langsung di variable atau parameter tanpa harus membuat function terlebih dahulu
	- Hal tersebut dinamakan anonymous function, atau function tanpa nama
*/

// Kode Program Anonymous Function
type Blacklist func(string) bool

func RegisterUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked ...")
	} else {
		fmt.Println("Welcome", name)
	}
}

func FunctionAnonymous() {
	blacklistSpam := func(name string) bool {
		return name == "anjing"
	}

	RegisterUser("anjing", blacklistSpam)
}
