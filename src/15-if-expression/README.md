# If Expression
- ```if``` adalah salah satu kata kunci yang digunakan untuk percabangan.
- Percabangan artinya kita bisa mengeksekusi kode program tertentu ketika suatu kondisi terpenuhi.
- Hampir di semua bahasa pemrograman mendukung if expression.

# Kode Program If
```go
package main

import "fmt"

func main() {
  name := "Sandy"

  if name == "Sandy" {
    fmt.Println("Hello Sandy")
  }
}
```

# Else Expression
- Blok ```if``` akan dieksekusi ketika bernilai true.
- Kadang kita ingin melakukan eksekusi program tertentu jika kondisi ```if``` bernilai false.
- Hal ini bisa dilakukan menggunakan ```else``` expression.

# Kode Program Else
```go
package main

import "fmt"

func main() {
  name := "Sandy"

  if name == "Sandy" {
    fmt.Println("Hello Sandy")
  } else {
    fmt.Println("Hi, boleh kenalan ?")
  }
}
```

# Else If Expression
- Kadang dalam ```if```, kita butuh membuat beberapa kondisi
- Kasus seperti ini, kita bisa menggunakan else if expression.

# Kode Program Else If Expression
```go
package main

import "fmt"

func main() {
  name := "Sandy"

  if name == "Sandy" {
    fmt.Println("Hello Sandy")
  } else if name == "Dwi" {
    fmt.Println("Hello Dwi")
  } else {
    fmt.Println("Hi, boleh kenalan ?")
  }
}
```

# If dengan Short Statement
- If mendukung sort statement sebelum kondisi.
- Hal ini sangat cocok untuk membuat statement yang sederhana sebelum melakukan pengecekan terhadap kondisi.

# Kode Program If Short Statement
```go
package main

import "fmt"

func main() {
  name := "Sandy"

  if length := len(name); length > 5 {
    fmt.Println("Nama terlalu panjang")
  } else {
    fmt.Println("Nama sudah benar")
  }
}
```

