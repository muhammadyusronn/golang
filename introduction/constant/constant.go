package main

import "fmt"

func main() {
	// Perbedaan constant dan variable adalah constant valuenya tidak bisa di replace
	// Constant juga tidak memberikan error atau warning ketika constant tersebut tidak digunakan
	// Contohnya pada constant beratBadan
	// Multiple Constant
	const (
		firstName = "Muhammad Yusron"
		lastName  = "Hartoyo"
	)
	const tempatLahir = "Palembang"
	const tanggalLahir = "1997-08-17"
	const agama = "Islam"
	const tinggiBadan = "178CM"
	const beratBadan = "Cie Kepo!" // Constant yang tidak digunakan
	fmt.Println("====================================")
	fmt.Println("Name : ", firstName, " ", lastName)
	fmt.Println("TTL : ", tanggalLahir, ", ", tempatLahir)
	fmt.Println("Religion : ", agama)
	fmt.Println("Tinggi Badan : ", tinggiBadan)
	fmt.Println("====================================")
}
