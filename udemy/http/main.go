package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	//fmt.Printf("%+v", resp)

	bs := make([]byte, 99999) // make an empty byte with space for 99999 elements
	resp.Body.Read(bs)        // Write response body to a byte slice
	fmt.Println(string(bs))   // Convert byte slice to string to be human readable

	io.Copy(os.Stdout, resp.Body) // Copy will write to some dest (Stdout) from a src (resp.Body)
}
