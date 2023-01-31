package src

import (
	"fmt"
	"sort"
)

/**
Package Sort
- Pakcage sort adalah package yang berisikan utilitas untuk proses pengurutan
- Agar data kita bisa diurutkan, kita harus mengimplementasikan kontrak di interface sort.interface
- https://golang.org/pkg/sort/
*/

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (value UserSlice) Len() int {
	return len(value)
}

func (value UserSlice) Less(i, j int) bool {
	return value[i].Age < value[j].Age
}

func (value UserSlice) Swap(i, j int) {
	value[i], value[j] = value[j], value[i]
}

func PackageSort() {
	users := []User{
		{"Sandy", 30},
		{"Dwi", 35},
		{"Handoko", 25},
		{"Trapsilo", 31},
	}

	sort.Sort(UserSlice(users))
	fmt.Println(users)
}
