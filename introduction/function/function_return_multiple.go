package main

import "fmt"

func person() (string, string, byte) {
	return "Muhammad Yusron Hartoyo", "Programmer", 25
}
func main() {
	name, job, age := person()
	fmt.Println("Name :", name, "Job :", job, "Age :", age)
	// Menghiraukan return
	names, jobs, _ := person()
	fmt.Println("Name :", names, "Job :", jobs)
}
