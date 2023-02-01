package src

import (
	"fmt"
	"regexp"
)

/**
Package Regexp
- Package regexp adalah utilitas Go-Lang untuk melakukan pencarian regular expression
- Regular expression di Go-Lang menggunakan library C yang dibuat Google bernama RE2
- https://github.com/google/re2/wiki/Syntax
- https://golang.org/pkg/regexp/
*/

/**
Beberapa Function di Package Regexp
- regexp.MustCompile(string) -> Membuat regexp
- Regexp.MatchString(string) bool -> Mengecek apakah Regexp match dengan string
- Regexp.FindAllString(string, max) -> Mencari string yang match dengan maximum jumlah hasil
*/

func PackageRegexp() {
	regex := regexp.MustCompile(`e([a-z])o`)

	fmt.Println(regex.MatchString("Sandy"))
	fmt.Println(regex.FindAllString("Sandy Andy Budi", 10))
}
