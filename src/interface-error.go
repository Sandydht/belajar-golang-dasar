package src

import (
	"errors"
	"fmt"
)

/**
Error Interface
- Go-Lang memiliki interface yang digunakan sebagai kontrak untuk membuat error, nama interface nya adalah error
*/

/**
Membuat Error
- Untuk membuat error, kita tidak perlu manual
- Go-Lang sudah menyediakan library untuk membuat helper secara mudah, yang terdapat di package errors (Package akan dibahas secara detail di materi tersendiri)
*/

// Kode Program Error Interface (1)
func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh NOL")
	} else {
		return nilai / pembagi, nil
	}
}

func InterfaceError() {
	result, err := Pembagian(10, 0)

	if err == nil {
		fmt.Println("Hasil:", result)
	} else {
		fmt.Println("Error:", err.Error())
	}
}
