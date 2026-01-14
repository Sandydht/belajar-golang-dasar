package main

import "fmt"

func main() {
	// Kode Program Switch
	name := "Sandy"

	switch name {
	case "Sandy":
		fmt.Println("Hello Sandy")
	case "Dwi":
		fmt.Println("Hello Dwi")
	default:
		fmt.Println("Hello what is your name ?")
	}

	// Kode Program Switch dengan Short Statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	// Kode Program Switch Tanpa Kondisi
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
