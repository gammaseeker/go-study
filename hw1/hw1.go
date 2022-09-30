package main

import (
	"fmt"
	"os"
) 

func main() {
	allArgs := os.Args
	indexedArgs := os.Args[1:]

	arg := os.Args[3]
	fmt.Println(allArgs)
	fmt.Println(indexedArgs)
	fmt.Println(arg)
}