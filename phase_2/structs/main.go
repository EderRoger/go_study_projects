package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	var p person
	p.firstName = "Alex"

	eder := person{firstName: "Eder Roger", lastName: "de Souza"}

	fmt.Println(eder)
	fmt.Printf("%+v", p)
}
