# Membuat Custom Error
- Karena error adalah sebuah interface, jadi jika kita ingin membuat error sendiri, kita bisa membuat ```struct``` yang mengikuti kontrak dari interface error.

# Kode Program Custom Error
```go
type validationError struct {
  Message string
}

func (v *validationError) Error() string {
  return v.Message
}

type notFoundError struct {
  Message string
}

func (v *notFoundError) Error() string {
  return v.Message
}
```

# Kode Program Menggunakan Custom Error
```go
func SaveData(id string, data any) error {
  if id == "" {
    return &validationError{Message: "validation error"}
  }

  if id != "sandy" {
    return &notFoundError{Message: "data not found"}
  }

  return nil
}
```

# Kode Program Mengecek Jenis Error
```go
package main

func main() {
  err := SaveData("", nil)

  if err != nil {
    if validationErr, ok := err.(*validationError); ok {
      fmt.Println("validation error", validationErr.Message)
    } else if notFoundErr, ok := err.(*notFoundError); ok {
      fmt.Println("not found error", notFoundErr.Message)
    } else {
      fmt.Println("unknown error", err.Error())
    }
  }
}
```