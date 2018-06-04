package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
}

func main() {

	alex := person{firstName: "Eder", lastName: "Roger"}

	fmt.Println(alex)
}
