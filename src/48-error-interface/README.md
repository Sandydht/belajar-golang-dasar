# Error Interface
- Go-Lang memiliki interface yang digunakan sebagai kontrak untuk membuat error, nama interface-nya adalah ```error```.
```go
type error interface {
  Error() string
}
```

# Membuat Error
- Untuk membuat error, kita tidak perlu manual.
- Go-Lang sudah menyediakan library untuk membuat helper secara mudah, yang terdapat di package errors (package akan dibahas secara detail di materi tersendiri).

# Kode Program Error Interface
```go
package main

import (
  "errors",
  "fmt"
)

func Pembagian(nilai int, pembagi int) (int, error) {
  if pembagi == 0 {
    return 0, errors.New("Pembagian dengan nol")
  } else {
    return nilai / pembagi, nil
  }
}

func main() {
  hasil, err := Pembagian(100, 10)

  if err == nil {
    fmt.Println("Hasil", hasil)
  } else {
    fmt.Println("Error", err.Error())
  }
}
```