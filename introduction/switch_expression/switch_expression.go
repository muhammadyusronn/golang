package main

import "fmt"

func main() {
	/*
	   Switch expression digunakan untuk melakukan pengecekan yang sederhana. Biasanya untuk mengecek satu variable
	*/
	name := "Firstda"
	switch name {
	case "Yusron":
		fmt.Println("Hallo, Yusron")
	case "Firstda":
		fmt.Println("Hello, Firstda")
	default:
		fmt.Println("Hah, lu siapa?")
	}

	// Switch dengan short statement

	switch length := len(name); length < 5 {
	case true:
		fmt.Println("Ukuran nama sudah pas")
	case false:
		fmt.Println("Nama terlalu panjang")
	}

	umur := 25

	switch {
	case umur > 30:
		fmt.Println("Usia anda udah dewasa")
	case umur < 30:
		fmt.Println("Usia anda masih muda")
	}
}
