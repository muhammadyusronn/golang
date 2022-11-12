package main

import "fmt"

/**
Struct merupakan template data yang digunakan untuk menggabungkan beberapa tipe data menjadi stau kesatuan
Data pada struct di simpan ke dalam field
Sederhananya, struct merupakan kumpulan dari beberapa field
*/
// Mendeklarasikan struct biasanya struct orang mendefinisikan diawali uppercase
type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var cust Customer
	cust.Name = "Muhammad Yusron Hartoyo"
	cust.Address = "Palembang"
	cust.Age = 25
	fmt.Println(cust)
	fmt.Println("STRUCT LITERALS")
	joko := Customer{
		Name:    "Joko",
		Address: "Jakarta",
		Age:     30,
	}
	budi := Customer{"Budi", "Bandung", 28} // Harus sesuai urutan field yang sudah di deklarasikan
	fmt.Println(joko)
	fmt.Println(budi)
}
