package main

import "fmt"

func main() {
	// Kode Program For
	counter := 1
	for counter <= 10 {
		fmt.Println("Perulangan ke", counter)
		counter++
	}

	// Kode Program For dengan Statement
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke", counter)
	}

	// Kode Program For Range
	names := []string{"Sandy", "Dwi", "Handoko", "Trapsilo"}
	for index, name := range names {
		fmt.Println("Perulangan ke", index, "=", name)
	}

	person := make(map[string]string)
	person["Name"] = "Sandy"
	person["TItle"] = "Programmer"
	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
