package main

import "fmt"

func main() {
	/*
	   Operasi perbandingan merupakan operasi yang membandingka dua buah data
	   Kembalian dari operasi perbandingan adalah boolean (true atau false)
	   Kembalian berupa true apabila hasil operasinya benar
	   Kembalian berupa false apabila hasil operasinya salah

	   Operato perbandingan
	   Operator					Keterangan
	   >						Lebih dari
	   <						Kurang dari
	   >=						Lebih dari sama dengan
	   <=						Kurang dari sama dengan
	   ==						Sama dengan
	   !=						Tidak sama dengan

	*/

	var name1 = "Yusron"
	var name2 = "Firstda"

	var result bool = name1 == name2
	fmt.Println(result)

	var value1 = 100
	var value2 = -100
	fmt.Println(value1 > value2)
}
