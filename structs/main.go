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
	contact   contactInfo
}

func main() {
	eder := person{
		firstName: "Eder",
		lastName:  "Roger",
		contact: contactInfo{
			email:   "e.r.s@gmail.com",
			zipCode: 34324454,
		},
	}

	fmt.Printf("%+v", eder)

}
