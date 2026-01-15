# Panic
- Panic function adalah function yang bisa kita gunakan untuk menghentikan program.
- Panic function biasanya dipanggil ketika terjadi panic pada saat program kita berjalan.
- Saat panic function dipanggil, program akan terhenti, namun defer function tetap akan dieksekusi.

# Kode Program Panic
```go
package main

import "fmt"

func endApp() {
  fmt.Println("End app")
}

func runApp(error bool) {
  defer endApp()
  if error {
    panic("ERROR")
  }
}

func main() {
  runApp(true)
}
```