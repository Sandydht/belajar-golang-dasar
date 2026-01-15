package main

import "fmt"

type Man struct {
	Name string
}

func (man Man) married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	sandy := Man{"Sandy"}
	fmt.Println(sandy.Name)
}
