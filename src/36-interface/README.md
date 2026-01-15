# Interface
- Interface adalah tipe data abstract, dia tidak memiliki implementasi langsung.
- Sebuah interface berisikan definisi-definisi method.
- Biasanya interface digunakan sebagai kontrak.

# Kode Program Interface
```go
import "fmt"

type HasName interface {
  getName() string
}

func sayHello(hasName HasName) {
  fmt.Println("Hello", hasName.getName())
}
```

# Implementasi Interface
- Setiap tipe data yang sesuai dengan kontrak interface, secara otomatis dianggap sebagai interface tersebut.
- Sehingga kita tidak perlu mengimplementasikan interface secara manual.
- Hal ini agak berbeda dengan bahasa pemrograman lain yang ketika membuat interface, kita harus menyebutkan secara eksplisit akan menggunakan interface mana.

# Kode Program Implementasi Interface (1)
```go
package main

type Person struct {
  Name string
}

func (person Person) getName() string {
  return person.Name
}

type HasName interface {
  getName() string
}

func sayHello(hasName HasName) {
  fmt.Println("Hello", hasName.getName())
}

func main() {
  person := Person{Name: "Sandy"}
  sayHello(person)
}
```

# Kode Program Implementasi Interface (2)
```go
package main

type Animal struct {
  Name string
}

func (animal Animal) getName() string {
  return animal.Name
}

type HasName interface {
  getName() string
}

func sayHello(hasName HasName) {
  fmt.Println("Hello", hasName.getName())
}

func main() {
  animal := Animal{Name: "Kucing"}
  sayHello(animal)
}
```