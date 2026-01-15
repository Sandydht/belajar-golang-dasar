# Function Return Value
- Function itu bisa mengembalikan data.
- Untuk memberitahu bahwa function mengembalikan data, kita harus menuliskan tipe data kembalian dari function tersebut.
- Jika function tersebut kita deklarasikan dengan tipe data pengembalian, maka wajib di dalam function-nya kita harus mengembalikan data.
- Untuk mengembalikan data dari function, kita bisa menggunakan kata kunci ```return```, diikuti dengan data-nya.

# Kode Program Function Return Value
```go
package main

import "fmt"

func getHello(name string) string {
  if name == "" {
    return "Hello bro"
  } else {
    return "Hello " + name
  }
}

func main() {
  result := getHello("Sandy")
  fmt.Println(result)
}
```