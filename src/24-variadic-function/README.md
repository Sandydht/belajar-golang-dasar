# Variadic Function
- Parameter yang berada di posisi terakhir, memiliki kemampuan dijadikan sebuah ```varargs```.
- ```varargs``` artinya data bisa menerima lebih dari satu input, atau anggap saja semacam array.
- Apa bedanya parameter biasa dengan tipe data array ?
  - Jika parameter tipe array, kita wajib membuat array terlebih dahulu sebelum mengirimkan ke function.
  - Jika parameter menggunakan ```varargs```, kita bisa langsung mengirim data-nya, jika lebih dari satu, cukup gunakan tanda koma.

# Kode Program Variadic Function
```go
package main

import "fmt"

func sumAll(numbers ...int) int {
  total := 0

  for _, number := range numbers {
    total += number
  }

  return total
}

func main() {
  total := sumAll(10, 10, 10, 10, 10)
  fmt.Println(total)
}
```

# Slice Paramater 
- Kadang ada kasus dimana kita menggunakan variadic function namun memiliki variable berupa slice.
- Kita bisa menjadikan slice sebagai ```varargs``` parameter.

# Kode Program Slice Parameter
```go
package main

import "fmt"

func sumAll(numbers ...int) int {
  total := 0

  for _, number := range numbers {
    total += number
  }

  return total
}

func main() {
  numbers := []int{20, 20, 20, 20, 20}
  total := sumAll(numbers...)
  fmt.Println(total)
}
```