package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	eder := person{
		firstName: "Eder Roger",
		lastName:  "de Souza",
		contactInfo: contactInfo{
			email:   "eder.r@gmail.com",
			zipCode: 909374937,
		},
	}
	eder.updateName("Ed")
	eder.print()

}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
