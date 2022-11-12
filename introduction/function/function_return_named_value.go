package main

import "fmt"

func getFullName() (firstName, middleName, lastName string) {
	firstName = "Muhamad"
	middleName = "Yusron"
	lastName = "Hartoyo"
	return
}
func main() {
	fullName, _, lastName := getFullName()
	fmt.Println(fullName)
	fmt.Println(lastName)
}
