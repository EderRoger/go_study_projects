package main

import "fmt"

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human  // anonymous field
	school string
}

type Employee struct {
	Human
	company string
}

// define a method in Human
func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

func main() {
	sam := Employee{Human{"Sam", 45, "55 83737474"}, "Go lang INC"}
	mark := Student{Human{"Mark", 25, "55 984735"}, "MIT"}

	sam.SayHi()
	mark.SayHi()
}
