package main

import "fmt"

type BlackList func(string) bool

func registerUser(name string, blacklist BlackList) {
	if blacklist(name) {
		fmt.Println("Welcome", name)
	} else {
		fmt.Println("You are blocked")
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "Yusron"
	}
	registerUser("Yusron", blacklist)
	registerUser("Firstda", func(name string) bool {
		return name == "Yusron"
	})
}
