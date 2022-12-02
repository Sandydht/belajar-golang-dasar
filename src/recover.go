package src

import "fmt"

/**
Recover
- Recover adalah function yang bisa kita gunakan untuk menangkap data panic
- Dengan recover proses panic akan terhenti
*/

// Kode Program Recover
func EndApplication() {
	message := recover()
	fmt.Println("Error dengan message:", message)
	fmt.Println("Aplikasi Selesai")
}

func RunningApplication(error bool) {
	defer EndApplication()
	if error {
		panic("APLIKASI ERROR")
	}
	fmt.Println("Aplikasi Berjalan")
}

func Recover() {
	RunningApplication(true)
	fmt.Println("Sandy")
}
