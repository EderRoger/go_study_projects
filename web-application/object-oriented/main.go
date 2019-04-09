package main

import "fmt"

type Rectangle struct {
	width, height float64
}

func area(r Rectangle) float64 {
	return r.width * r.height
}

func main() {
	r1 := Rectangle{10, 40}
	r2 := Rectangle{20, 10}

	fmt.Println("Area 1 is: ", area(r1))
	fmt.Println("Area 2 is: ", area(r2))

}
