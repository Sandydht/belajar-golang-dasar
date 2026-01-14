# Tipe Data Map
- Pada array atau slice, untuk mengakses data, kita menggunakan index number dimulai dari 0.
- Map adalah tipe data lain yang berisikan kumpulan data yang sama, namun kita bisa menentukan jenis tipe data index yang akan kita gunakan.
- Sederhananya, map adalah tipe data kumpulan key-value (kata kunci-nilai), dimana kata kuncinya bersifat unik, tidak boleh sama.
- Berbeda dengan array dan slice, jumlah data yang kita masukkan ke dalam map boleh sebanyak-banyaknya, asalkan kata kunci-nya berbeda, jika kita gunakan kata kunci sama, maka secara otomatis data sebelumnya akan diganti dengan data baru.

# Kode Program Map
```go
package main

import "fmt"

func main() {
  person := map[string]string{
    "name": "Sandy",
    "address": "Jakarta"
  }

  fmt.Println(person)
  fmt.Println(person["name"])
  fmt.Println(person["address"])
}
```

# Function Map
| Operasi | Keterangan |
|---------|------------|
| len(map) | Untuk mendapatkan jumlah data di map |
| map[key] | Mengambil data di map dengan key |
| map[key] = value | Mengubah data di map dengan key |
| make(map[TypeKey]TypeValue) | Membuat map baru |
| delete(map, key) | Menghapus data di map dengan key |

# Kode Program Map
```go
package main

import "fmt"

func main() {
  book := make(map[string]string)
  book["title"] = "Buku Go-Lang"
  book["author"] = "Sandy Dwi Handoko Trapsilo"
  book["wrong"] = "Ups"

  delete(book, "wrong")

  fmt.Println(book)
}
```
