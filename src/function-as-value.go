package src

import "fmt"

/**
Function Value
- Function adalah first class citizen
- Function juga merupakan tipe data dan bisa di simpan di dalam variable
*/

// Kode Program Function Value
func GetGoodBye(name string) string {
	return "Good Bye " + name
}

func FunctionAsValue() {
	goodbye := GetGoodBye
	fmt.Println(goodbye("Sandy"))
}
