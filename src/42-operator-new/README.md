# Operator New
- Sebelumnya untuk membuat pointer dengan menggunakan operator ```&```.
- Go-Lang juga memiliki function ```new``` yang bisa digunakan untuk membuat pointer.
- Namun function ```new``` mengembalikan pointer ke data kosong, artinya tidak ada data awal.

# Kode Program Function New 
```go
package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
  alamat1 := new(Address)
  alamat2 := alamat1

  alamat2.Country = "United Kingdom"

  fmt.Println(alamat1) // alamat1 berubah
  fmt.Println(alamat2)
}
```