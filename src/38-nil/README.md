# Nil
- Biasanya di dalam bahasa pemrograman lain, object yang belum diinisialisasi maka secara otomatis nilai-nya adalah null atau nil.
- Berbeda dengan Go-Lang, di Go-Lang saat kita buat variable dengan tipe data tertentu, maka secara otomatis akan dibuatkan default value-nya.
- Namun di Go-Lang ada data nil, yaitu data kosong.
- Nil sendiri hanya bisa digunakan di beberapa tipe data, seperti interface, function, map, slice, pointer, dan channel.

# Kode Program Nil
```go
package main

import "fmt"

func newMap(name string) map[string]string {
  if name == "" {
    return nil
  } else {
    return map[string]string{
      "name": name
    }
  }
}

func main() {
  data := newMap("")

  if data == nil {
    fmt.Println("Data kosong")
  } else {
    fmt.Println(data)
  }
}
```