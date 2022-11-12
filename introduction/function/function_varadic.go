package main

import "fmt"

/*
Variable argument harus berada diposisi paling belakang dan hanya bisa satu dalam satu function
*/

func sumAll(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}

func main() {
	total := sumAll(10, 10, 10, 10, 10)
	fmt.Println(total)
	fmt.Println("Variadic function dengan menggunakan slice")
	numbers := []int{10, 10, 10, 50}
	total = sumAll(numbers...)
	fmt.Println(total)
}
