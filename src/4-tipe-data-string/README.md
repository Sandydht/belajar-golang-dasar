# Tipe Data String
- String adalah kumpulan karakter.
- Jumlah karakter di dalam string bisa nol sampai tidak terhingga.
- Tipe data string di Go-Lang direpresentasikan dengan kata kunci string.
- Nilai data string di Go-Lang selalu diawali dengan karakter **" (petik dua)** dan diakhiri dengan karakter **" (petik dua)**.

# Kode Program String
```go
package main

import "fmt"

func main() {
  fmt.Println("Sandy")
  fmt.Println("Sandy Dwi")
  fmt.Println("Sandy Dwi Handoko")
  fmt.Println("Sandy Dwi Handoko Trapsilo")
}
```

# Function untuk String
| Function | Keterangan |
|----------|------------|
| len("string") | Menghitung jumlah karakter di String |
| "string"[number] | Mengambil karakter pada posisi yang ditentukan |

# Kode Program String
```go
package main

import "fmt"

func main() {
  fmt.Println(len("Sandy"))
  fmt.Println("Sandy Dwi"[0])
  fmt.Println("Sandy Dwi Handoko"[1])
  fmt.Println("Sandy Dwi Handoko Trapsilo"[2])
}
```