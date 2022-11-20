package src

import "fmt"

/*
	For
	- Dalam bahasa pemrograman, biasanya ada fitur yang bernama perulangan
	- Di Go-Lang fitur perulangan hanya ada satu, yaitu for
*/

/*
	For dengan Statement
	- Dalam for, kita bisa menambahkan statement, dimana terdapat 2 statement yang bisa ditambahkan di for
	- Init statement, yaitu statement sebelum for di eksekusi
	- Post statement, yaitu statement yang akan selalu di eksekusi di akhir tiap perulangan
*/

/*
	For Range
	- For bisa digunakan untuk melakukan iterasi terhadap semua data collection
	- Data collection contohnya array, slice, dan map
*/

func For() {
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
