package main

import "fmt"

/*
DEFER
Defer merupakan function yang bisa dijadwalkan untuk diekseskusi setelah sebuah function selesai dijalankan
Defer function akan selalu dijalankan walaupun terjadi error ketika dieksekusi

PANIC
Panic function merupakan function digunakan untuk memberhentikan program
Ketika panic function dipanggil, program akan berhenti, namun defer function akan tetap jalan

RECOVER
Recover merupakan function yang dapat digunakan untuk mengambil data dari panic
Dengan recover proses panic akan berhenti, dan aplikasi akan tetap berjalan
*/

func deferFunction() {
	fmt.Println("Memanggil defer function")
}

func runFunc(value int) {
	defer deferFunction() // Memanggil defer function
	fmt.Println("Running Application")
	result := 10 / value
	fmt.Println("Result : ", result)
}

// Implementasi Panic function
func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("Terjadi error : ", message)
	}
	fmt.Println("APLIKASI BERAKHIR")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("APLIKASI ERROR!")
	}
	fmt.Println("APLIKASI BERJALAN")
}
func main() {
	fmt.Println("IMPELENTASI DEFER FUNCTION")
	runFunc(10) // successfully running
	// runFunc(0)  // Ketika function di running terjadi error karena 10 / 0 maka defer function tetap dijalankan
	fmt.Println("IMPELENTASI PANIC FUNCTION DAN RECOVER FUNCTION")
	runApp(true)
	fmt.Println("TES APLIKASI RUNNING") // Dengan recover function line 52 akan tetap di running
}
