package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	var p person
	p.firstName = "Alex"

	eder := person{
		firstName: "Eder Roger",
		lastName:  "de Souza",
		contact: contactInfo{
			email:   "eder.r@gmail.com",
			zipCode: 909374937,
		},
	}

	fmt.Println(eder)
	fmt.Printf("%+v", p)
}
