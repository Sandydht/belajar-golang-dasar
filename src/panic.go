package src

import "fmt"

/**
Panic
- Panic function adalah function yang bisa kita gunakan untuk menghentikan program
- Panic function biasanya dipanggil ketika terjadi error pada saat program kita berjalan
- Saat panic function dipanggil, program akan terhenti, namun defer function tetap akan dieksekusi
*/

// Kode Program Panic
func EndApp() {
	fmt.Println("Aplikasi Selesai")
}

func RunApp(error bool) {
	defer EndApp()
	if error {
		panic("APLIKASI ERROR")
	}
}

func Panic() {
	RunApp(false)
}
