package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(`Exercise 001: Write a program which will 
	find all such numbers which are divisible by 7 
	but are not a multiple of 5, between 2000 and 3200 (both included).  
	The numbers obtained should be printed in a comma-separated sequence 
	on a single line.`)

	res := ex001(2000, 3200)
	fmt.Println(res)
}

func ex001(low, high int) string {
	var numbers []string
	for i := low; i<= high; i++ {
		if i%7 == 0 {
			if i%5 == 0 {
				continue
			} else {
				numbers = append(numbers, strconv.Itoa(i))
			}
		}

		/*
		if i%7 == 0 && i%5 != 0 {
			numbers = append(numbers, strconv.Iota(i))
		}
		*/
	}
	return strings.Join(numbers, ",")
}