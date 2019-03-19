package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	eder := person{firstName: "Eder Roger", lastName: "de Souza"}
	fmt.Println(eder)
}
