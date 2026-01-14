package main

import "fmt"

func main() {
	// Kode Program Map (1)
	person := map[string]string{
		"name":    "Sandy",
		"address": "Kendal",
	}

	person["title"] = "Programmer"

	fmt.Println("person =", person)
	fmt.Println("name =", person["name"])
	fmt.Println("address =", person["address"])

	// Kode Program Map (2)
	var book map[string]string = make(map[string]string)

	book["title"] = "Belajar Golang Dasar"
	book["author"] = "Sandy"
	book["ups"] = "ups"

	fmt.Println("book =", book)
	delete(book, "ups")
	fmt.Println("book =", book)
	fmt.Println("len(book) =", len(book))
}
