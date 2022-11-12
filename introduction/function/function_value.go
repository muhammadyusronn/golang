package main

import "fmt"

func getGoodBye(name string) string {
	return "Goodbye, " + name
}
func main() {
	// Menyimpang sebuah function ke dalam variable
	sayGoodBye := getGoodBye("Guys!")
	fmt.Println(sayGoodBye)
}
