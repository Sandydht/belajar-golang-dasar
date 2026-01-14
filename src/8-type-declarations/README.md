# Type Declarations
- Type declarations adalah kemampuan membuat ulang tipe data baru dari tipe data yang sudah ada.
- Type declarations biasanya digunakan untuk membuat alias terhadap tipe data yang sudah ada, dengan tujuan agar lebih mudah dimengerti.

# Kode Program Type Declarations
```go
package main

import "fmt"

func main() {
  type NoKTP string

  var ktpSandy NoKTP = "1111111111"
  fmt.Println(ktpSandy)
  fmt.Println(NoKTP("2222222222"))
}
```