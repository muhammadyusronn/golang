package main

import "fmt"

func main() {
	// Penggunaan looping di golang hanya menggunakan for loops

	counter := 1

	for counter <= 10 {
		fmt.Println("Pengulangan ke-", counter)
		counter++
	}

	// For Loops dengan statement
	for i := 1; i <= 10; i++ {
		fmt.Println("Pengulangan dengan statement ke-", i)
	}

	// Pengulangan dengan slice

	slice := []string{"Muhammad Yusron Hartoyo", "Firstda Agustin", "Rosa Alawiyah"}
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
	// Menampilkan index dari slice (For Range)
	for i, value := range slice {
		fmt.Println("Index", i, "=", value)
	}

	// Hanya menampilkan value tidak menampilkan index
	for _, value := range slice {
		fmt.Println(value)
	}

	// Pengulangan dengan map

	person := make(map[string]string)
	person["name"] = "Muhammad Yusron Hartoyo"
	person["job"] = "Programmer"

	// Menampilkan key dan value
	for key, value := range person {
		fmt.Println(key, "=", value)
	}

	// Menampilkan value
	for _, value := range person {
		fmt.Println(value)
	}
}
