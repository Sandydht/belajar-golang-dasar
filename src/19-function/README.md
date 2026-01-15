# Function
- Sebelumnya kita sudah mengenal sebuah function yang wajib dibuat agar program kita bisa berjalan, yaitu function ```main```.
- Function adalah sebuah blok kode yang sengaja dibuat dalam program agar bisa digunakan berulang-ulang.
- Cara membuat function-nya sangat sederhana, hanya dengan menggunakan kata kunci ```func``` lalu diikuti dengan nama function-nya dan blok kode isi function-nya.
- Setelah membuat function, kita bisa mengeksekusi function tersebut dengan memanggilnya menggunakan kata kunci nama function-nya diikuti tanda kurung buka, kurung tutup.

# Kode Program Function
```go
package main 

import "fmt"

func sayHello() {
  fmt.Println("Hello, world!")
}

func main() {
  sayHello()
}
```