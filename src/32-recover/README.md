# Recover
- Recover adalah function yang bisa kita gunakan untuk menangkap data panic.
- Dengan recover proses panic akan terhenti, sehingga program akan tetap berjalan.

# Kode Program Recover Salah
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

  message := recover()
  fmt.Println("Terjadi error", message)
}

func main() {
  runApp(true)
}
```

# Kode Program Recover Benar
```go
package main

import "fmt"

func endApp() {
  fmt.Println("End app")
  
  message := recover()
  fmt.Println("Terjadi error", message)
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