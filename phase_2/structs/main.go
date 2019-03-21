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

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (pointerToPerson person) print() {
	fmt.Printf("%+v", pointerToPerson)
}
