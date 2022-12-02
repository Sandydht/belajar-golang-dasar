package src

import "fmt"

/**
Defer
- Defer function adalah function yang bisa kita jadwalkan untuk dieksekusi setelah sebuah function selesai di eksekusi
- Defer function akan selalu dieksekusi walaupun terjadi error di function yang dieksekusi
*/

// Kode Program Defer
func Logging() {
	fmt.Println("Selesai memanggil function")
}

func RunApplication(value int) {
	defer Logging()
	fmt.Println("Run Application")
	result := 10 / value
	fmt.Println("Result: ", result)
}

func Defer() {
	RunApplication(10)
}
