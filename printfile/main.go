package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	if len(os.Args) == 1 {
		fmt.Println("File name expected as a parameter Example: \"go run main.go filename\"")
		os.Exit(1)
	}

	fileName := os.Args[1]

	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Error to open file:", fileName, "error", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, file)
}
