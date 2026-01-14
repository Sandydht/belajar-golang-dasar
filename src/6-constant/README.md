# Constant
- Constant adalah variable yang nilai-nya tidak bisa diubah lagi setelah pertama kali diberi nilai.
- Cara pembuatan constant mirip dengan variable, yang membedakan hanya kata kunci yang digunakan adalah ```const```, bukan ```var```.
- Saat pembuatan ```constant```, kita wajib langsung menginisialisasikan data-nya.

# Kode Program Constant
```go
package main

import "fmt"

func main() {
  const firstName string = "Sandy"
  const lastName = "Dwi"

  // error
  firstName = "Tidak bisa diubah"
  lastName = "Tidak bisa diubah"

  fmt.Println(firstName)
  fmt.Println(lastName)
}
```

# Deklarasi Multiple Constant
- Sama seperti variable, di Go-Lang kita juga bisa membuat multiple constant.

# Kode Program Constant
```go
package main

import "fmt"

func main() {
  const (
    firstName = "Sandy"
    lastName = "Dwi"
  )

  // error
  firstName = "Tidak bisa diubah"
  lastName = "Tidak bisa diubah"

  fmt.Println(firstName)
  fmt.Println(lastName)
}
```