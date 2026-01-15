package main

import (
	"belajar-golang-dasar/src/helper"
	"belajar-golang-dasar/src/other"
	"fmt"
)

func main() {
	helperSayHello := helper.SayHello("Sandy")
	fmt.Println(helperSayHello)

	otherSayHello := other.SayHello("Sandy")
	fmt.Println(otherSayHello)
}
