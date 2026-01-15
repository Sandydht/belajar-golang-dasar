# Namen Return Values
- Biasanya saat kita memberi tahu bahwa sebuah function mengembalikan value, maka kita hanya mendeklarasikan tipe data return value di function.
- Namun kita juga membuat variable secara langsung di tipe data return function-nya.

# Kode Program Namen Return Values
```go
package main

import "fmt"

func getCompleteName() (firstName string, middleName string, lastName string) {
  firstName = "Sandy"
  middleName = "Dwi"
  lastName = "Handoko"

  return firstName, middleName, lastName
}

func main() {
  firstName, middleName, lastName := getCompleteName()
	fmt.Println(firstName, middleName, lastName)
}
```