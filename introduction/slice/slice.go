package main

import "fmt"

func main() {
	/*
		   Tipe data slice adalah potongan dari array
		   Perbedaannya slice panjang atau ukurannya bisa berubah

		   Detail tipe data slice
		   Tipe data slice ada 3 = pointer, length, dan capacity
		   Pointer -> Petunjuk data pertama di array slice
		   Length -> Panjang slice
		   Capacity -> Kapasitas dari slice, length tidak boleh lebih dari capacity

		   Membuat Slice dari array
		   Membuat Slice				Keterangan
		   array[low:high]				Membuat slice dari array dimulai index low sampai index sebelum high
		   array[low:]					Membuat slice dari array dimulai dari index low sampai index akhir di array
		   array[:high]					Membuat slice dari array dimual dari index 0 sampai index sebelum high
		   array[:]						Membuat slice dari array dimulai dari index 0 sampei index akhir di array


		   Function Slice
		   Function								Keterangan
		   len(slice)							Menghitung panjang slice
		   cap(slice)							Menghitung kapasitas dari slice
		   append(slice, data)					Membuat slice baru dengan menambahkan data ke posisi terakhir slice,
		   										Jika kapasitas sudah penuh, maka akan membuat array baru
			make([]TypeData, length, capacity)	Membuat slice baru
			copy(destination, source)			Menyalin slice dari source ke destination

	*/

	months := [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	fmt.Println("Default array : ", months) // Array yang sudah didefinisikan

	slice1 := months[3:6]                   // Membuat slice dari index ke-3 sampai index ke-6 dari array months
	fmt.Println("Slice Pertama : ", slice1) // Menampilkan slice yang sudah dibuat
	/*
		Merubah data pada slice di posisi index ke 2 (Juni)
		Perubahan yang terjadi mengakibatkan perubahan value pada array months
		karena sebenarnya perubahan pada slice tetap disimpan pada array months
	*/
	slice1[2] = "Bulan Juni"
	fmt.Println(slice1)

	slice2 := months[10:]                 // Membuat slice kedua dari index ke-10 sampai index paling terakhir dari array months
	fmt.Println("Slice kedua : ", slice2) // Mencetak slice kedua
	/*
		Membuat slice ke-3 dengan menambahkan bulan baru dari value slice2
		Kapasitas dari slice ke-2 adalah 2 sedangkan ketika kita menambahkan bulan baru maka sudah melebihi kapasitas yang ada
		sehingga akan secara otomatis membuat array baru yang mengakibatkan tidak terjadi perubahan pada array months
	*/
	slice3 := append(slice2, "Bulan Baru")

	fmt.Println(slice3) // Mencetak slice ketika (otomatis membuat array baru)
	fmt.Println(months) // Mencetak array months untuk mengecek apakah ada perubaha value yang diakibatkan dari pembuatan slice3

	// Membuat Slice
	newSlice := make([]string, 3, 5) // Membuat slice dengan panjang slice adalah 3 dan  kapasitas slice adalah 5

	newSlice[0] = "Muhammad Yusron Hartoyo"
	newSlice[1] = "Firstda Agustin"
	newSlice[2] = "Rosa Alawiyah"

	fmt.Println(newSlice)      // Mencetas slice
	fmt.Println(len(newSlice)) // Menghitung panjang slice
	fmt.Println(cap(newSlice)) // Mendapatkan kapasitas slice

	// Menyalin slice (copy(Destination, Source))
	// Sebelum menyalin slice pastikan panjang slice dan kapasitas slice minimal sama dengan slice yang akan di copy
	// Jika ukuran slice baru lebih kecil dari slice yang akan dicopy makan datanya akan terpotong seperti line 93
	terpotongSlice := make([]string, 1, cap(newSlice))
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)                                      // Menyalin slice newSlice ke copySlice
	copy(terpotongSlice, newSlice)                                 // Menyalin slice newSlice ke terpotongSlice
	fmt.Println("Contoh slice yang ukurannya benar : ", copySlice) // Menampilkan selice yang baru di salin ke copySlice
	fmt.Println("Contoh slice terpotong : ", terpotongSlice)       // Menampilkan selice yang baru di salin ke terpotongSlice

	// Hati - hati dalam mendeklarasikan slice dan array
	iniArray := [5]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}
	fmt.Println("ini adalah array : ", iniArray)
	fmt.Println("ini adalah slice : ", iniSlice)

}
