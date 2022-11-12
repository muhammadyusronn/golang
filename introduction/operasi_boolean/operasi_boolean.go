package main

import "fmt"

func main() {
	/*
		   		Operasi Boolean
		Operator			Keterangan
		&&					Dan
		||					Atau
		!					Kebalikan

	*/
	var name = "Muhammad Yusron Hartoyo"
	var nilaiAkhir = 85
	var nilaiAbsen = 83

	var lulusNilaiAkhir = nilaiAkhir > 80
	var lulusNilaiAbsen = nilaiAbsen > 80

	var lulus bool = lulusNilaiAbsen && lulusNilaiAkhir

	fmt.Println("Apakah", name, "lulus?", lulus)

}
