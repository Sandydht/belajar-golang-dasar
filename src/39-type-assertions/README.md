# Type Assertions
- Type assertions merupakan kemampuan merubah tipe data menjadi tipe data yang diinginkan.
- Fitur ini sering sekali digunakan ketika kita bertemu dengan data interface kosong.

# Kode Program Type Assertions
```go
package main 

import "fmt"

func random() interface{} {
  return "OK"
}

func main() {
  result := random()
  resultString := result.(string)
  fmt.Println(resultString)

  resultInt := result.(int) // panic
  fmt.Println(resultInt)
}
```

# Type Assertions Menggunakan Switch
- Saat salah menggunakan type assertions, maka bisa berakibat terjadi ```panic``` di aplikasi kita.
- Jika ```panic``` tidak ter-recover, maka otomatis program kita akan mati.
- Agar lebih aman, sebaiknya kita menggunakan switch expression untuk melakukan type assertions.

# Kode Program Type Assertions Switch
```go
package main

import "fmt"

func random() interface{} {
  return "OK"
}

func main() {
  result := random()

  switch value := result.(type) {
    case string:
      fmt.Println("String", value)
    case int:
      fmt.Println("Int", value)
    default:
      fmt.Println("Unknown")
  }
}
```