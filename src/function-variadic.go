package src

import "fmt"

/**
Variadic Function
- Parameter yang berada di posisi terakhir, memiliki kemampuan dijadikan sebuah varargs
- Varargs artinya datanya bisa menerima lebih dari satu input, atau anggap saja semacam array
- Apa bedanya dengan parameter biasa dengan tipe data array ?
  - Jika parameter tipe array, kita wajib membuat array terlebih dahulu sebelum mengirimkan ke function
	- Jika parameter menggunakan varargs, kita bisa langsung mengirim data nya, jika lebih dari satu, cukup gunakan tanda koma
*/

/**
Slice Parameter
- Kadang ada kasus dimana kita menggunakan Variadic Function, namun memiliki variable berupa slice
- Kita bisa menjadikan slice sebagai varargs parameter
*/

// Kode Program Variadic Function
func SumAll(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func FunctionVariadic() {
	total := SumAll(10, 10, 10, 10, 10)
	fmt.Println(total)

	numbers := []int{10, 10, 10, 10, 10}
	total = SumAll(numbers...)
	fmt.Println(total)
}
