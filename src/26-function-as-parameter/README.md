# Function as Parameter
- Function tidak hanya bisa kita simpan di dalam variable sebagai value.
- Namun juga bisa kita gunakan sebagai parameter untuk function lain.

# Kode Program Function as Parameter
```go
package main

import "fmt"

func sayHelloWithFilter(name string, filter func(string) string) {
  fmt.Println("Hello ", filter(name))
}

func spamFilter(name string) string {
  if name == "Anjing" {
    return "..."
  } else {
    return name
  }
}

func main() {
  sayHelloWithFilter("world", spamFilter)

  filter := spamFilter
  sayHelloWithFilter("Anjing", filter)
}
```

# Function Type Declarations
- Kadang jika function terlalu panjang, agak ribet untuk menuliskannya di dalam parameter.
- Type declarations juga bisa digunakan untuk membuat alias function, sehingga akan mempermudah kita menggunakan function sebagai parameter.

# Kode Program Type Declarations
```go
package main

import "fmt"

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
  fmt.Println("Hello ", filter(name))
}

func spamFilter(name string) string {
  if name == "Anjing" {
    return "..."
  } else {
    return name
  }
}

func main() {
  sayHelloWithFilter("world", spamFilter)

  filter := spamFilter
  sayHelloWithFilter("Anjing", filter)
}
```