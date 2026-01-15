# Break & Continue
- Break & Continue adalah kata kunci yang bisa digunakan dalam perulangan.
- Break digunakan untuk menghentikan seluruh perulangan.
- Continue digunakan untuk menghentikan perulangan yang berjalan, dan langsung melanjutkan ke perulangan selanjutnya.

# Kode Program Break
```go
package main

import "fmt"

func main() {
  for i := 0; i < 10; i++ {
    if i == 5 {
      break
    }

    fmt.Println("Perulangan ke =", i)
  }
}
```

# Kode Program Continue
```go
package main

import "fmt"

func main() {
  for i := 0; i < 10; i++ {
    if i % 2 == 0 {
      continue
    }

    fmt.Println("Perulangan ke =", i)
  }
}
```