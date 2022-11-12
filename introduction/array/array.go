package main

import "fmt"

func main() {
	/*
		Kumpulan data yang memiliki tipe data yang sama
		Daya tampung array tidak bisa bertambah setelah array dibuat


		Index

		index					Data
		0						Muhamamd Yusron Hartoyo
		1						Firstda Augustin
		2						Rosa Alawiyah

		Function array
		Function				Keterangan
		len(array)				Untuk menghitung jumlah array atau panjang array
		array[index]			Untuk mendapatkan posisi array beradasarkan index
		array[index]			Untuk mengubah value di posisi index
	*/

	// Cara mendeklarasikan array
	var names [3]string
	// Mengisi array nama
	names[0] = "Muhammad Yusron Hartoyo"
	names[1] = "Firstda Agustin"
	names[2] = "Rosa Alawiyah"

	// Membuat dan mendeklarasikan array usia
	var ages = [3]byte{
		25, 27, 23,
	}
	fmt.Println("Banyak orang : ", len(names))
	fmt.Println("============================================")
	fmt.Println(names[0], ages[0], "th")
	fmt.Println("============================================")
	fmt.Println("============================================")
	fmt.Println(names[1], ages[1], "th")
	fmt.Println("============================================")
	fmt.Println("============================================")
	fmt.Println(names[2], ages[2], "th")
	fmt.Println("============================================")
}
