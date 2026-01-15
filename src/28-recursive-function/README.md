# Recursive Function
- Recursive function adalah function yang memanggil function dirinya sendiri.
- Kadang dalam pekerjaan, kita sering menemui kasus dimana menggunakan recursive function lebih mudah.
- Contoh masus yang lebih mudah diselesaikan menggunakan recursive adalah factorial.

# Kode Program Factorial For Loop
```go
package main

import "fmt"

func factorialLoop(value int) int {
  result := 1

  for i := value; i > 0; i-- {
    result *= i
  } 

  return result
}

func main() {
  fmt.Println(factorialLoop(10))
}
```

# Kode Program Factorial Recursive
```go
package main

import "fmt"

func factorialRecursive(int value) int {
  if value == 1 {
    return 1
  } else {
    return value * factorialRecursive(value-1)
  }
}

func main() {
  fmt.Println(factorialRecursive(10))
}
```