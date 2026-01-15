package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked ...")
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	blacklistSpam := func(name string) bool {
		return name == "anjing"
	}

	registerUser("John", blacklistSpam)
	registerUser("anjing", func(name string) bool {
		return name == "anjing"
	})
}
