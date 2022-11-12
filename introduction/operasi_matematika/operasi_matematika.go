package main

import "fmt"

func main() {
	var result = 10 + 10
	fmt.Println(result)

	// Bisa juga dengan menggunakan variable
	var a = 10
	var b = 10
	var c = a * b
	fmt.Println(c)

	// a = a + 1
	var i = 10
	i += 10 // i = i + 10
	fmt.Println(i)
	i++            // i = i + 1
	fmt.Println(i) // i = i + 1
}
