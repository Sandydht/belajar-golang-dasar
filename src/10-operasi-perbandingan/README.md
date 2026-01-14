# Operasi Perbandingan
- Operasi perbandingan adalah operasi untuk membandingkan dua buah data.
- Operasi perbandingan adalah operasi yang menghasilkan nilai boolean (benar atau salah).
- Jika hasil operasi-nya adalah benar, maka nilainya adalah true.
- Jika hasil operasi-nya adalah salah, maka nilainya adalah false.

| Operator | Keterangan |
|----------|------------|
| > | Lebih dari |
| < | Kurang dari |
| >= | Lebih dari sama dengan |
| <= | Kurang dari sama dengan |
| == | Sama dengan |
| != | Tidak sama dengan |

# Kode Program Operasi Perbandingan
```go
package main

import "fmt"

func main() {
  var name1 = "Sandy"
  var name2 = "Dwi"

  var result bool = name1 == name2

  fmt.Println(result)
}
```