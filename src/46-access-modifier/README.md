# Access Modifier
- Di bahasa pemrograman lain, biasanya ada kata kunci yang bisa digunakan untuk menentukan access modifier terhadap suatu function atau variable.
- Di Go-Lang, untuk menentukan access modifier, cukup dengan nama function atau variable.
- Jika nama-nya diawali dengan huruf besar, maka artinya bisa diakses dari package lain, jika dimulai dengan huruf kecil, artinya tidak bisa diakses dari package lain.

# Kode Program Access Modifier
```go
package accessmodifier

var version = "1.0.0" // tidak bisa diakses dari luar package
var Application = "golang"

func sayGoodBye(name string) string { // tidak bisa diakses dari luar package
  return "Good bye " + name
}

func SayHello(name string) string {
  return "Hello " + name
}
```