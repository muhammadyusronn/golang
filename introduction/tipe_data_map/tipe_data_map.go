package main

import "fmt"

func main() {
	/*
	   Map adalah tipe data lain yang  berisikan data yang sama, namun kita bisa menentuka tipe data index yang akan kita gunakan
	   Map adalah tipe data kumpulan key-value (kata kunci - nilai), dimana kata kuncinya bersifat unique, tidak boleh sama
	   Map memiliki perbedaan dengan slice dan array karena map bisa memasukan data sebanyak mungkin tanpa ada batasan
	   dengan syarat key atau kata kuncinya bersifat unik, jika key atau kata kuncinya sama maka secara otomatis
	   akan mengubah data pada key sebelumnya

	   // Function
	   Function							Keterangan
	   len(map)							Menghitung jumlah data pada map
	   map[key]							Mengambil data di map dengan menggunakan key
	   map[key] = value					Mengubah data di map dengan menggunakan key
	   make(map[TypeKey]TypeValue)		Membuat map baru
	   delete(map,key)					Menghapus data di map dengan menggunakan key

	*/

	// Cara mendeklarasikan Map

	person := map[string]string{
		"name":    "Muhammad Yusron Hartoyo",
		"address": "Jalan Otto Iskandar Dinatta",
	}
	person["job"] = "Programmer"                // Menambahkan key baru pada map person
	fmt.Println(person)                         // Menampilkan map person
	fmt.Println("Nama : ", person["name"])      // Mengakses nama dari person
	fmt.Println("Alamat : ", person["address"]) // Mengakses alamat dari person
	fmt.Println("Pekerjaan : ", person["job"])  // Mengakses alamat dari person

	// Membuat map baru
	book := make(map[string]string)
	book["title"] = "Buku pemrograman Go-Lang"
	book["author"] = "Muhammad Yusron Hartoyo"
	book["wrong"] = "Ups"
	fmt.Println("Map sebelum book[wrong] dihapus : ", book)
	delete(book, "wrong") // Menghapus data map yang memiliki key wrong
	fmt.Println(book)     // Mencetak map book (key-value dari wrong akan terhapus)
}
