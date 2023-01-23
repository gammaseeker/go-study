package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func main() {
	fmt.Println("Hello please enter a starting boundary: ")
	var response string
	fmt.Scan(&response)
	start, err := strconv.Atoi(response)
	if err != nil {
		fmt.Println("Please enter a valid integer")
		return
	}
	fmt.Println("User typed: ", start)
	fmt.Println("Please enter an ending boundary: ")
	fmt.Scan(&response)
	end, err := strconv.Atoi(response)
	if end <= start {
		fmt.Println("Please enter an ending boundary that is greater than the starting")
		return
	}

	if err != nil {
		fmt.Println("Please enter a valid integer")
		return
	}
	fmt.Println("User typed: ", end)

	randomTarget := rand.Intn(end) + start
	fmt.Println("A random number has been generated. Please guess the number:")
	var guess int = -1
	for guess != randomTarget {
		fmt.Scan(&response)
		guess, err = strconv.Atoi(response)
		if err != nil {
			fmt.Println("Please enter a valid integer")
			continue
		}
		if guess == randomTarget {
			fmt.Println("Correct")
			return
		}
		fmt.Println("Please try again")
	}
}
