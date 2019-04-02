package main

import "fmt"

func main() {
	s := "hello"
	c := []byte(s)
	c[0] = 's'
	s2 := string(c)

	fmt.Println(s2)
}
