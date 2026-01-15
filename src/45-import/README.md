# Import 
- Secara standar, file Go-Lang hanya bisa mengakses file Go-Lang lainnya yang berada dalam package yang sama.
- Jika kita ingin mengakses file Go-Lang yang berada di luar package, maka kita bisa menggunakan ```import```.

# Kode Program Import
```go
package main

import (
  "belajar-golang-dasar/helper",
  "fmt"
)

func main() {
  result := helper.SayHello("Sandy")
  fmt.Printlin(result)
}
```