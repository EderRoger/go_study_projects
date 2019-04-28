package main

import (
	"fmt"
	"os"
)

type isInt func(int) bool

func isOdd(number int) bool {
	return number%2 != 0
}

func isEven(number int) bool {
	return number%2 == 0
}

func filter(slice []int, f isInt) []int {
	var result []int
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}

var slice = []int{1, 2, 3, 4, 5, 6, 7}

// structs

type Skills []string

type Human struct {
	name   string
	age    int
	weight int
}

type Student struct {
	Human
	Skills
	int
	specialty string
}

func main() {
	skills := []string{"fly", "strength"}
	eder := Student{Human: Human{name: "Eder", age: 34, weight: 78}, Skills: skills, specialty: "IT"}
	fmt.Println(eder)
	// odd := filter(slice, isOdd)
	// even := filter(slice, isEven)

	// fmt.Println("Odd: ", odd)
	// fmt.Println("Even: ", even)

	// s := "hello"
	// c := []byte(s)
	// c[0] = 's'
	// s2 := string(c)

	// fmt.Println(s2)

	// import "fmt"
	// import "os"

	// const i = 100
	// const pi = 3.1415
	// const prefix = "Go_"

	// var i int
	// var pi float32
	// var prefix string

	// Group form

	// import (
	// 	"fmt"
	// 	"os"
	// )

	// const (
	// 	i = 100
	// 	pi = 3.1415
	// 	prefix = "Go_"
	// )

	// var (
	// 	i int
	// 	pi float32
	// 	prefix string
	// )
	// goTO()
}

func goTO() {
	i := 0

Here:
	fmt.Println(i)
	i++
	if i == 100 {
		goto Out
	}

	goto Here
Out:
	os.Exit(1)
}
