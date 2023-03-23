package main

import (
	"fmt"
	"sync"
)

// ALWAYS pass wg as a pointer, never as a value
func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done() // defer says "do not execute this line of code until after the current function exits"
	// wg.Done() decrements wg by 1
	fmt.Println(s)
}

func main() {
	var wg sync.WaitGroup

	//go printSomething("This is the first thing to be printed!")
	words := []string{
		"alpha",
		"beta",
		"gamma",
		"delta",
		"epsilion",
		"zeta",
	}
	//wg.Add(6) //Wait for 6 things to finish before proceeding
	//wg.Add(999) If ctr is more than number of goroutines being spawned, get deadlock error that all goroutines are asleep
	wg.Add(len(words))

	for i, x := range words {
		go printSomething(fmt.Sprintf("%d: %s", i, x), &wg)
		// go routines do not execute in order of which they spawn. Execute based on what go scheduler decides
	}
	wg.Wait() // Wait for wg to reach 0

	//time.Sleep(1 * time.Second) // This is the bad way of implementing waiting

	wg.Add(1) // Never let wg go below 0
	printSomething("This is the second thing to be printed!", &wg)
}
