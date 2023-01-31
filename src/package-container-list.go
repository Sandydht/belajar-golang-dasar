package src

import (
	"container/list"
	"fmt"
)

/**
Package container/list
- Package container/list adalah implementasi struktur data double linked list di Go-Lang
- https://golang.org/pkg/container/list
*/

func PackageContainerList() {
	data := list.New()
	data.PushBack("Sandy")
	data.PushBack("Dwi")

	fmt.Println("data front:", data.Front().Value)
	fmt.Println("data back:", data.Back().Value)

	// Loop dari depan ke belakang
	// for element := data.Front(); element != nil; element = element.Next() {
	// 	fmt.Println("element:", element.Value)
	// }

	// Loop dari belakang ke depan
	for element := data.Back(); element != nil; element = element.Prev() {
		fmt.Println("element:", element.Value)
	}
}
