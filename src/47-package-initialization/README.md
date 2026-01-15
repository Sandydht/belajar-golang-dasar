# Package Initialization
- Saat kita membuat package, kita bisa membuat function yang akan diakses ketika package kita diakses.
- Ini sangat cocok ketika package kita berisi function-function untuk berkomunikasi dengan database, kita membuat function inisialisasi untuk membuka koneksi ke database.
- Untuk membuat function yang diakses secara otomatis ketika package diakses, kita cukup membuat function dengan nama init.

# Kode Program Initialization (1)
```go
package database

var connection string

func init() {
  connection = "MySQL"
}

func GetDatabase() string {
  return connection
}
```

# Kode Program Initialization (2)
```go
package main

import (
	"belajar-golang-dasar/src/database"
	"fmt"
)

func main() {
  fmt.Println(database.GetDatabase())
}
```