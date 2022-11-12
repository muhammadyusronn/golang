package main

import "fmt"

func main() {
	// Cara mendeklarasikan variable
	// Cara pertama
	var name string
	name = "Muhammad Yusron Hartoyo"
	// Cara kedua (Langsung memberikan nilainya)
	var usia = 25
	// Cara Ketiga
	country := "Indonesia"
	// Cara ke-empat deklarasi variable secara multiple
	var (
		alamat    = "Jalan Otto Iskanda Dinata"
		kecamatan = "Jatinegara"
		kota      = "Jakarta Timur"
		provinsi  = "DKI Jakarta"
	)
	fmt.Println("")
	fmt.Println("=========================================")
	fmt.Println("")
	fmt.Println("Name : ", name)
	fmt.Println("Age :", usia)
	fmt.Println("Country :", country)
	fmt.Println("Address : ", alamat, " Kecamatan : ", kecamatan, " Kota : ", kota, " Provinsi : ", provinsi)
	name = "Firstda Agustin"
	usia = 27
	country = "Republik Indonesia"
	fmt.Println("")
	fmt.Println("=========================================")
	fmt.Println("")
	fmt.Println("Replace name :", name)
	fmt.Println("Replace age :", usia)
	fmt.Println("Replace country :", country)
	fmt.Println("Address : ", alamat, " Kecamatan : ", kecamatan, " Kota : ", kota, " Provinsi : ", provinsi)
	fmt.Println("")
	fmt.Println("=========================================")
	fmt.Println("")

}
