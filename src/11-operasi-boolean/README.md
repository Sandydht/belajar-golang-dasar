# Operasi Boolean
| Operator | Keterangan |
|----------|------------|
| && | Dan |
| \|\| | Atau |
| ! | Kebalikan |

# Operasi &&
| Nilai 1 | Operator | Nilai 2 | Hasil |
|---------|----------|---------|-------|
| true | && | true | true |
| true | && | false | false |
| false | && | true | false |
| false | && | false | false |

# Operasi ||
| Nilai 1 | Operator | Nilai 2 | Hasil |
|---------|----------|---------|-------|
| true | \|\| | true | true |
| true | \|\| | false | true |
| false | \|\| | true | true |
| false | \|\| | false | false |

# Operasi !
| Operator | nilai | Hasil |
|----------|-------|-------|
| ! | true | false |
| ! | false | true |

# Kode Program Operasi Boolean
```go
package main

import "fmt"

func main() {
  var nilaiAkhir = 90 
  var absensi = 80

  var lulusNilaiAkhir bool = nilaiAkhir > 80
  var lulusAbsensi bool = absensi > 80

  var lulus bool = lulusNilaiAkhir && lulusAbsensi

  fmt.Println(lulus)
}
```