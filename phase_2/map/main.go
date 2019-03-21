package main

import "fmt"

func main() {

	colors := map[string]string{
		"red":   "#FFF000",
		"green": "#F1F002",
		"white": "#FFFFFF",
	}

	printMap(colors)
}

func printMap(colors map[string]string) {
	for color, value := range colors {
		fmt.Printf("Hex code for %v is %v ", color, value)
	}
}
