package main

import "fmt"

func main() {
	/*
		Kata kunci yang digunakan dalam perulangan
		Break digunakan untuk menghentikan perulangan
		Continue digunakan untuk menghentikan perulangan berjalan dan lanjut ke perulangan berikutnya
	*/
	// Perulangan yang hanya menampilkan data apabila bernilai genap
	for i := 0; i <= 20; i++ {
		if i%2 != 0 {
			continue
		}
		// Apabila i bernilai > 10 maka perulangan dihentikan dengan break
		if i > 10 {
			break
		}
		fmt.Println("Nilai :", i)
	}
}
