package src

import "fmt"

/*
	Recursive Function
	- Recursive function adalah function yang memanggil dirinya sendiri
	- Kadang dalam pekerjaan, kita sering menemui kasus yang memerlukan recursive function
	- Contoh kasus yang lebih mudah diselesaikan menggunakan recursive function adalah Factorial
*/

// Kode Program Factorial For Loop
func FactorialLoop(value int) int {
	result := 1

	for i := value; i > 0; i-- {
		result *= i
	}

	return result
}

// Kode Program Factorial Recursive
func FactorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * FactorialRecursive(value-1)
	}
}

func FunctionRecursive() {
	fmt.Println(FactorialLoop(5))
	fmt.Println(FactorialRecursive(5))
}
