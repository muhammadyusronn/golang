package main

import "fmt"

// Mendeklarasikan tipe data function untuk filter
type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Halowwwww," + nameFiltered)
}

func filter(name string) string {
	if name == "anjing" {
		return "*****"
	} else {
		return name
	}
}
func main() {
	sayHelloWithFilter("Yusron", filter)
}
