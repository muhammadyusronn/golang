package main

import "fmt"

func main() {
	/*
	   If adalah salah satu kata kunci yang digunakan untuk percabangan di golang
	   Percabangan artinya suatu kode akan dieksekusi apabila memenuhi suatu kondisi tertentu
	*/
	name := "M. Ricky Hidayat"
	if name == "Yusron" {
		fmt.Println("Hallo, ", name)
	} else if name == "Firstda" {
		fmt.Println("Hallo, ", name)
	} else {
		fmt.Println("Woi lu siapa sih ", name, "?")
	}

	// Go-Lang bisa juga short statement
	if length := len(name); length <= 5 {
		fmt.Println("Nama anda sudah sesuai")
	} else {
		fmt.Println("Nama anda terlalu panjang")
	}
}
