# For
- Dalam bahasa pemrograman, biasanya ada fitur yang bernama perulangan.
- Salah satu fitur perulangan adalah for loops.

# Kode Program For
```go
package main

import "fmt"

func main() {
  counter := 1

  for counter <= 10 {
    fmt.Println("Perulangan ke =", counter)
    counter++
  }
}
```

# For dengan Statement
- Dalam for, kita bisa menambahkan statement, dimana terdapat 2 statement yang bisa ditambahkan.
- Init statement, yaitu statement sebelum for dieksekusi.
- Post statement, yaitu statement yang akan selalu dieksekusi di akhir tiap perulangan.

# Kode Program For dengan Statement
```go
package main

import "fmt"

func main() {
  for counter := 1; counter <= 10; counter++ {
    fmt.Println("Perulangan ke =", counter)
  }
}
```

# For Range
- For bisa digunakan untuk melakukan iterasi terhadap semua data collection.
- Data collection contohnya array, slice, dan map.

# Kode Program For Range
```go
package main

import "fmt"

func main() {
  names := []string{"Sandy", "Dwi", "Handoko", "Trapsilo"}

  for index, name := range names {
    fmt.Println("index", index, "=", name)
  }
}
```