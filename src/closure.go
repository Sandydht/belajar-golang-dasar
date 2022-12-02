package src

import "fmt"

/**
Closures
- Closure adalah kemampuan sebuah function berinteraksi dengan data - data disekitarnya dalam scope yang sama
- Harap gunakan fitur closure ini dengan bijak saat kita membuat aplikasi
*/

// Kode Program Closure
func Closure() {
	counter := 0
	name := "Sandy"

	increment := func() {
		fmt.Println("Increment")
		counter++
		name := "Dwi"
		fmt.Println(name)
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)
}
