package main

import "fmt"

/*
Parameter tidak waji
Parameter dari function bisa lebih dari 1 dengan syarat nama parameter harus unik
Apabila function memiliki parameter, maka pada saat memanggil function wajib memasukan parameternya
*/
func sayHallo(name string) {
	fmt.Println("Halo,", name)
}
func main() {
	sayHallo("Yusron")
}
