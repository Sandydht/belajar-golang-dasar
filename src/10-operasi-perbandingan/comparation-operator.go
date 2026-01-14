package main

import "fmt"

func main() {
	// Kode Program Operasi Perbandingan
	var name1 string = "Sandy"
	var name2 string = "Dwi"
	var result bool = name1 == name2
	fmt.Println("result =", result)

	var value1 = 100
	var value2 = 200
	fmt.Println("value1 > value2 =", value1 > value2)
	fmt.Println("value1 < value2 =", value1 < value2)
	fmt.Println("value1 == value2 =", value1 == value2)
	fmt.Println("value1 != value2 =", value1 != value2)
}
