package main

import (
	"fmt"
	"os"
)

func main() {
	s := "hello"
	c := []byte(s)
	c[0] = 's'
	s2 := string(c)

	fmt.Println(s2)

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
	goTO()
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
