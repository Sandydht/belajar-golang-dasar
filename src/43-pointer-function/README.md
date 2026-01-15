# Pointer Method
- Walaupun method akan menempel di ```struct```, tapi sebenarnya data ```struct``` yang diakses di dalam method adalah pass by value.
- Sangat direkomendasikan menggunakan pointer di method, sehingga tidak boros memory karena harus selalu duplikasi ketika memanggil method.

# Kode Program Pointer di Method (1)
```go
package main

type Man struct {
  Name string
}

func (man Man) married() {
  man.Name = "Mr. " + man.Name
}

func main() {
  sandy := Man{"Sandy"}
  sandy.married()
}
```

# Kode Program Pointer di Method (2)
```go
package main

import "fmt"

type Man struct {
  Name string
}

func (man *Man) married() {
  man.Name = "Mr. " + man.Name
}

func main() {
  sandy := Man{"Sandy"}
  sandy.married()
}
```