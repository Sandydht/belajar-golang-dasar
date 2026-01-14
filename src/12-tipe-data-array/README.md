# Tipe Data Array
- Array adalah tipe data yang berisikan kumpulan data dengan tipe yang sama.
- Saat membuat array, kita perlu menentukan jumlah data yang bisa ditampung oleh Array tersebut.
- Daya tampung array tidak bisa bertambah setelah Array dibuat.

# Index di Array
| Index | Data |
|-------|------|
| 0 | Sandy |
| 1 | Dwi |
| 2 | Handoko |
| 3 | Trapsilo |

# Kode Program Array
```go
package main

import "fmt"

func main() {
  var names [4]string
  names[0] = "Sandy"
  names[1] = "Dwi"
  names[2] = "Handoko"
  names[3] = "Trapsilo"

  fmt.Println(names[0])
  fmt.Println(names[1])
  fmt.Println(names[2])
  fmt.Println(names[3])
}
```

# Membuat Array Langsung
- Di Go-Lang kita juga bisa membuat Array secara langsung saat deklarasi variable.

# Kode Program Array
```go
package main

import "fmt"

func main() {
  var values = [int]{90, 80, 95}
  
  fmt.Println(values)
}
```

# Function Array
| Operasi | Keterangan |
|---------|------------|
| len(array) | Untuk mendapatkan panjang array |
| array[index] | Mendapatkan data di posisi index |
| array[index] = value | Mengubah data di posisi index |

# Kode Program Array
```go
package main

import "fmt"

func main() {
  var values = [...]int{90, 80, 95}

  fmt.Println(values)
  fmt.Println(len(values))
  values[0] = 100
  fmt.Println(value)
}
```

