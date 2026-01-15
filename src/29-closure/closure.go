package main

import "fmt"

func main() {
	counter := 0
	name := "Sandy"

	increment := func() {
		fmt.Println("Increment")
		counter++
		name := "Dwi"
		fmt.Println(name)
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)
}
