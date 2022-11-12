package main

import "fmt"

func main() {
	counter := 1

	increment := func() {
		counter := 2
		a := "variabel dalam function(a) hanya bisa diakses di dalam function sedangkan variabel diluar function (counter) bisa diakses di dalam function"
		fmt.Println(a)
		fmt.Println("Variabel dalam function (counter):", counter)
		counter++
	}
	increment()
	fmt.Println(counter)
}
