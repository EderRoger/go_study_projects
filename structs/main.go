package main

import (
	"fmt"
)

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
		firstName: "Eder",
		lastName:  "Roger",
		contactInfo: contactInfo{
			email:   "e.r.s@gmail.com",
			zipCode: 34324454,
		},
	}

	ederPointer := &eder

	ederPointer.updateName("Bacano")
	eder.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
