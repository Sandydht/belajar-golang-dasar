package main

import "fmt"

func main() {
	// Kode Program Type Declarations
	type NoKTP string
	type Married bool

	var ktpSandy NoKTP = "1111111111"
	var marriedStatus Married = true
	fmt.Println(ktpSandy)
	fmt.Println(NoKTP("2222222222"))
	fmt.Println(marriedStatus)
}
