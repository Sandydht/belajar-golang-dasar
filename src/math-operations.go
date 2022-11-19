package src

import "fmt"

/*
	Operasi Matematika
	- + -> penjumlahan
	- - -> pengurangan
	- * -> perkalian
	- / -> pembagian
	- % -> sisa hasil bagi (modulo)
*/

/*
	Augmented Assignments
	- a = a + 10 -> a += 10
	- a = a - 10 -> a -= 10
	- a = a * 10 -> a *= 10
	- a = a / 10 -> a /= 10
	- a = a % 10 -> a %= 10
*/

/*
	Unary Operator
	- ++ -> a = a + 1
	- -- -> a = a - 1
	- - -> negative
	- + -> positive
	- ! -> boolean kebalikan
*/

func MathOperations() {
	// Kode Program Operasi Matematika (1)
	var a = 10
	var b = 10
	var c = a + b
	fmt.Println(c)

	// Kode Program Operasi Matematika (2)
	var i = 10
	i += 10
	fmt.Println(i)

	// Kode Program Operasi Matematika (3)
	i++
	fmt.Println(i)

	var negative = -100
	var positive = +100 // by default sudah positif
	fmt.Println(negative)
	fmt.Println(positive)
}
