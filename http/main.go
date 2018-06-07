package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// byteSlice := make([]byte, 99999)
	// resp.Body.Read(byteSlice)
	// fmt.Println(string(byteSlice))

	lw := logWriter{}

	io.Copy(lw, resp.Body) //advanced form to print bytes

}

func (logWriter) Write(byteSlice []byte) (int, error) {
	fmt.Println(string(byteSlice))
	fmt.Println("Just wrote this many bytes:", len(byteSlice))

	return len(byteSlice), nil
}
