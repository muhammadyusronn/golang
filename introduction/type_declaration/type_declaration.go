package main

import "fmt"

func main() {
	/*
		Type Declaration adalah kemampuan untuk membuat tipe data baru dari tipe data yang sudah ada
		Tujuan dari pembuatan tipe data baru ini adalah untuk membuat alias terhadap tipe data yang sudah ada
		dengan tujuan agar mudah dimengerti
	*/
	type NoKTP string
	type married bool

	var NIK NoKTP = "12345678910"
	var marriedStatus married = true
	fmt.Println(NIK)
	fmt.Println(marriedStatus)

}
