package main

import "fmt"

type Person struct {
	Name, Adress string
	age          int
}

func (person Person) sayHello() {
	fmt.Println("Hallo, ", person.Name)
}
func main() {
	yusron := Person{Name: "Muhammad Yusron Hartoyo"}
	yusron.sayHello()
}
