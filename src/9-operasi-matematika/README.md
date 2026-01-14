# Operasi Matematika
| Operator | Keterangan |
|----------|------------|
| + | Penjumlahan |
| - | Pengurangan |
| * | Perkalian |
| / | Pembagian |
| % | Sisa hasil bagi |

# Kode Program Operasi Matematika
```go
package main

import "fmt"

func main() {
  var a = 10
  var b = 10
  var c = a + b

  fmt.Println(c)
}
```

# Augmented Assignments
| Operasi Matematika | Augmented Assignments |
|--------------------|-----------------------|
| a = a + 10 | a += 10 |
| a = a - 10 | a -= 10 |
| a = a * 10 | a *= 10 |
| a = a / 10 | a /= 10 |
| a = a % 10 | a %= 10 |

# Kode Program Augmented Assignments
```go
package main

import "fmt"

func main() {
  var i = 10
  i += 10

  fmt.Println(i)
}
```

# Unary Operator 
| Operator | Keterangan |
|----------|------------|
| ++ | a = a + 1 |
| -- | a = a - 1 |
| - | Negative |
| + | Positive |
| ! | Boolean kebalikan |

# Kode Program Unary Operator
```go
package main

import "fmt"

func main() {
  var j = 1 
  j++ // j = j + 1
  j++

  fmt.Println(j)
}
```