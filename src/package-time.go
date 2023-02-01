package src

import (
	"fmt"
	"time"
)

/**
Package time
- Package time adalah package yang berisikan fungsionalitas untuk management waktu di Go-Lang
- https://golang.org/pkg/time/
*/

/**
Beberapa Function di Package Time
- time.Now() -> Untuk mendapatkan waktu saat ini
- time.Date(...) -> Untuk membuat waktu
- time.Parse(layout, string) -> Untuk memparsing waktu dari siang
*/

func PackageTime() {
	now := time.Now()
	fmt.Println("now:", now)

	utc := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println("utc:", utc)

	parse, _ := time.Parse(time.RFC3339, "2006-01-02")
	fmt.Println("parse:", parse)
}
