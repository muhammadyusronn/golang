package main

import "fmt"

func main() {
	// Konversi tipe data dapat dilakukan dengan syarat tipe data yang baru memiliki ukuran yang cukup untuk menampung value pada tipe data yang sebelumnya
	var nilai32 int32 = 100000
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32) // Tipe data yang baru tidak cukup

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8) // Output akan menjadi - karena tipe data int8 tidak cukup menampung value nilai32 sehingga dia akan kembali ke titik yang paling bawah

	// Konversi byte ke string
	var name = "Firstda"
	var char = name[0]
	fmt.Println(name)         // Mencetak nama
	fmt.Println(char)         // Mencetak karakter pertama dari variabel name (sebelum dikonversi ke string)
	fmt.Println(string(char)) // Mengonversi value cari char ke string

}
