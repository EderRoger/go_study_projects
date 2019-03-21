package main

import "fmt"

func main() {

	colors := make(map[string]string)

	colors["white"] = "#fffff"

	delete(colors, "white")

	fmt.Println(colors)
}
