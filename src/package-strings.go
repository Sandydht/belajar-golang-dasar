package src

import (
	"fmt"
	"strings"
)

/**
Package String
- Package string adalah package yang berisikan function - function untuk memanipulasi tipe data String
- Ada banyak sekali function yang bisa kita gunakan
*/

/**
Beberapa Function di Package Strings
- strings, Trim(string, outset) -> Memotong outset di awal dan akhir string
- strings, ToLower(string) -> Membuat semua karakter string menjadi lower case
- strings, ToUpper(string) -> Membuat semua karakter string menjadi upper case
- strings, Split(string, separator) -> Memotong string berdasarkan separator
- strings, Contains(string, search) -> Mengecek apakah string mengandung string lain
- strings, ReplaceAll(string, from, to) -> Mengubah semua string dari from ke to
*/

func PackageStrings() {
	fmt.Println(strings.Contains("Sandy Dwi Handoko Trapsili", "Sandy"))
	fmt.Println(strings.Split("Sandy Dwi Handoko Trapsilo", " "))
	fmt.Println(strings.ToLower("SANDY DWI HANDOKO TRAPSILO"))
	fmt.Println(strings.ToUpper("sandy dwi handoko trapsilo"))
	fmt.Println(strings.ToTitle("sandy dwi handoko trapsilo"))
	fmt.Println(strings.Trim("    Sandy Dwi Handoko Trapsilo    ", " "))
	fmt.Println(strings.ReplaceAll("Sandy Dwi Handoko Trapsilo", "Sandy", "SND"))
}
