package main

import "fmt"

func main() {
	// Kode Program If
	name := "Sandy"

	if name == "Sandy" {
		fmt.Println("Hello Sandy")
	} else if name == "Dwi" {
		fmt.Println("Hello Dwi")
	} else {
		fmt.Println("Hello what is your name ?")
	}

	// Kode Program If Short Statement
	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
