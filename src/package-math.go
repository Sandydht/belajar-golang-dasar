package src

import (
	"fmt"
	"math"
)

/**
Package Math
- Package math merupakan package yang berisikan constant dan fungsi matematika
- https://golang.org/pkg/math/
*/

/**
Beberapa Function di Package Math
- math.Round(float64) -> Membulatkan float64 keatas atau kebawah, sesuai dengan yang paling dekat
- math.Floor(float64) -> Membulatkan float64 kebawah
- math.Ceil(float64) -> Membulatkan float64 keatas
- math.Max(float64, float64) -> Mengembalikan nilai float64 paling besar
- math.Min(float64, float64) -> Mengembalikan nilai float64 paling kecil
*/

func PackageMath() {
	fmt.Println("Round:", math.Round(3.56))
	fmt.Println("Floor:", math.Floor(3.56))
	fmt.Println("Ceil:", math.Ceil(3.14))
	fmt.Println("Max:", math.Max(2.14, 3.14))
	fmt.Println("Min:", math.Min(2.14, 3.14))
}
