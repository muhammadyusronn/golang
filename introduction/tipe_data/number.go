package main

import "fmt"

func main() {
	/*
		Tipe Data			Minimul Value			Maximum Value
		int8			    -127					127
		int16				-32.768					32.767
		int32				-2.147.483.648			2.147.483.647
		int64				-9223372036854775808	9223372036854775807
	*/

	/*
		Tipe Data			Minimul Value			Maximum Value
		uint8			    0						255
		uint16				0						65.535
		uint32				0						4.294.967.295
		uint64				0						18446744073709551615
	*/

	// Alias
	/*
		Tipe Data				Alias Untuk
		byte					uint8
		rune					int32
		int						minimal int32
		unint					minmal unit32
	*/
	fmt.Println("Satu = ", 1)
	fmt.Println("Dua = ", 2)
	fmt.Println("Tiga Koma 5 = ", 3.5)

}
