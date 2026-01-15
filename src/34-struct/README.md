# Struct 
- Struct adalah sebuah template data yang digunakan untuk menggabungkan nol atau lebih tipe data lainnya dalam satu kesatuan.
- Struct biasanya representasi data dalam program aplikasi yang kita buat.
- Data di struct disimpan dalam field.
- Sederhananya struct adalah kumpulan dari field.
 
# Kode Program Struct
```go
type Customer struct {
  Name, Address string
  Age           int
}
```

# Membuat Data Struct
- Struct adalah template data atau prototype data.
- Struct tidak bisa langsung digunakan.
- Namun kita bisa membuat data/object dari struct yang telah kita buat.

# Kode Program Membuat Data Struct
```go
package main

import "fmt"

type Customer struct {
  Name, Address string
  Age           int
}

func main() {
  var customer Customer
  customer.Name = "Sandy Dwi Handoko Trapsilo"
  customer.Address = "Jakarta, Indonesia"
  customer.Age = "28"

  fmt.Println(customer)
}
```

# Struct Literals
- Sebelumnya kita telah membuat data dari struct, namun sebenarnya ada banyak cara yang bisa kita gunakan untuk membuat data dari struct.

# Kode Program Struct Literals
```go
package main

import "fmt"

type Customer struct {
  Name, Address string
  Age           int
}

func main() {
  joko := Customer{
    Name: "Joko",
    Address: "Jakarta, Indonesia",
    Age: 30
  }
  fmt.Println(joko)

  budi := Customer{"Budi", "Jakarta, Indonesia", 30}
  fmt.Println(budi)
}
```