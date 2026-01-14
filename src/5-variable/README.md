# Variable
- Variable adalah tempat untuk menyimpan data.
- Variable digunakan agar kita bisa mengakses data yang sama dimanapun kita mau.
- Di Go-Lang variable hanya bisa menyimpan tipe data yang sama, jika kita ingin menyimpan data yang berbeda-beda jenis, kita harus membuat beberapa variable.
- Untuk membuat variable, kita bisa menggunakan kata kunci ```var```, lalu diikuti dengan nama variable dan tipe datanya.

# Kode Program Variable
```go
package main

import "fmt"

func main() {
  var name string

  name = "Sandy Dwi Handoko Trapsilo"
  fmt.Println(name)

  name = "Budi Setiawan"
  fmt.Println(name)
}
```

# Tipe Data Variable
- Saat kita membuat variable, maka kita wajib menyebutkan tipe data variable tersebut.
- Namun jika kita langsung menginisialisasikan data pada variable-nya, maka kita tidak wajib menyebutkan tipe data variable-nya.

# Kata Kunci Var
- Di Go-Lang, kata kunci ```var``` saat membuat variable tidak lah wajib.
- Asalkan saat membuat variable kita langsung menginisialisasi datanya.
- Agar tidak perlu menggunakan kata kunci ```var```, kita perlu menggunakan kata kunci ```:=``` saat menginisialisasikan data pada variable tersebut.

# Kode Program Variable
```go
package main

import "fmt"

func main() {
  name := "Sandy Dwi Handoko Trapsilo" // hanya untuk deklarasi awal saja
  fmt.Println(name)

  name = "Budi Setiawan"
  fmt.Println(name)
}
```

# Deklarasi Multiple Variable
- Di Go-Lang kita bisa membuat multiple variable.
- Code yang dibuat akan lebih bagus dan mudah dibaca.

# Kode Program Variable
```go
package main

import "fmt"

func main() {
  var (
    firstName = "Sandy"
    lastName = "Dwi"
  )

  fmt.Println(firstName)
  fmt.Println(lastName)
}
```