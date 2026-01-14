# Switch Expression
- Selain if expression, untuk melakukan percabangan, kita juga bisa menggunakan switch expression.
- Switch expression sangat sederhana dibandingkan if.
- Biasanya switch expression digunakan untuk melakukan pengecekan ke kondisi dalam satu variable.

# Kode Program Switch
```go
package main

import "fmt"

func main() {
  name := "Sandy"

  switch name {
    case "Sandy":
      fmt.Println("Hello Sandy")
    case "Dwi":
      fmt.Println("Hello Dwi")
    default:
      fmt.Println("Hi, boleh kenalan ?")
  }
}
```

# Switch dengan Short Statement
- Sama dengan if, switch juga mendukung short statement.

# Kode Program Switch dengan Short Statement
```go
package main

import "fmt"

func main() {
  name := "Sandy"

  switch length := len(name); length > 5 {
    case true:
      fmt.Println("Nama terlalu panjang")
    case false:
      fmt.Println("Nama sudah benar")
  }
}
```

# Switch Tanpa Kondisi
- Kondisi di switch expression tidak wajib.
- Jika kita menggunakan kondisi di switch expression, kita bisa menambahkan kondisi tersebut di setiap case-nya.

# Kode Program Switch Tanpa Kondisi
```go
package main

import "fmt"

func main() {
  name := "Sandy"
  length := len(name)

  switch {
    case length > 10:
      fmt.Println("Nama terlalu panjang")
    case length > 5:
      fmt.Println("Nama lumayan panjang")
    default:
      fmt.Println("Nama sudah benar")
  }
}
```

