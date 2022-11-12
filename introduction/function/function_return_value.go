package main

import "fmt"

/* Tipe data return function harus sama dengan tipe data yang di return */
func sayHello(name string) string {
	if name == "" {
		return "Hello, Bro!"
	} else {
		return "Hello," + name
	}
}
func main() {
	fmt.Println(sayHello(""))       // Return Hello, Bro!
	fmt.Println(sayHello("Yusron")) // Return Hello, Yusron!

}
