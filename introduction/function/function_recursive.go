package main

import "fmt"

/*
Recursive function merupakan function yang memanggil dirinya sendiri
Recursive function harus memiliki kondisi berhenti seperti line 20
Contoh kasus yang menggunakan recursive fuction adalah Factorial
*/
func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

func factorialRecursive(value int) int {
	// Kondisi berhenti
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}
func main() {
	input := factorialLoop(10)
	fmt.Println("Melakukan factorial tanpa function recursive", input)
	input = factorialRecursive(10)
	fmt.Println("Melakukan factorial dengan function recursive", input)
}
