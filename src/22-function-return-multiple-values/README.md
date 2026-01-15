# Returning Multiple Values
- Function tidak hanya dapat mengembalikan satu value, tapi juga bisa multiple values.
- Untuk memberitahu jika function mengembalikan multiple values, kita harus menulis semua tipe data return value-nya di function.

# Kode Program Returning Multiple Values
```go
package main

import "fmt"

func getFullName() (string, string, string) {
  return "Sandy", "Dwi", "Handoko"
}

func main() {
  firstName, middleName, _ := getFullName()
  fmt.Println(firstName, middleName)
}
```