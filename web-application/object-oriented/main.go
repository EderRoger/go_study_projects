package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (c Circle) Area() float64 {
	return c.radius * c.radius * math.Pi
}

func main() {
	r1 := Rectangle{10, 40}
	r2 := Rectangle{20, 10}
	c1 := Circle{30}

	fmt.Println("Area 1 is: ", r1.Area())
	fmt.Println("Area 2 is: ", r2.Area())
	fmt.Println("Circle 1 is: ", c1.Area())

}
