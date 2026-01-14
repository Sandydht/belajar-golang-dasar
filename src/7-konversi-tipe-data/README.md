# Konversi Tipe Data
- Di Go-Lang kadang kita butuh melakukan konversi suatu tipe data ke tipe lain.
- Misal kita ingin mengkonversi tipe data ```int32``` ke ```int64```, dan lain-lain.

# Kode Program Konversi Tipe Data (1)
```go
package main

import "fmt"

func main() {
  var nilai32 int32 = 32768
  var nilai64 int64 = int64(nilai32)
  var nilai16 int16 = int16(nilai32)

  fmt.Println(nilai32)
  fmt.Println(nilai64)
  fmt.Println(nilai16) // Bernilai negatif karena melebihi kapasitas (number overflow), oleh karena itu value nya akan kembali ke awal / nilai minimum (bisa dilihat table nya di materi 2-tipe-data-number).
}
```

# Kode Program Konversi Tipe Data (2)
```go
package main

import "fmt"

func main() {
  var name = "Sandy Dwi Handoko Trapsilo"
  var e = name[0] // Mengembalikan byte dari huruf "S"
  var eString = string(e)

  fmt.Println(name)
  fmt.Println(eString)
}
```