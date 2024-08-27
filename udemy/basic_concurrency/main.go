package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

/*
Concurrency is NOT parallelism

Concurrency - We can have multiple threads executing code. If one thread blocks, another one is picked up and worked on
  - "We can schedule work and switch between them on the fly"

Parallelism - Multiple threads executed at the exact same time. Requires multiple CPUs
  - "We can do multiple things at the exact same time"

Go routine is created every time we run a go program (main go routine).
Go routine is like an "engine that runs code."
Go routine executes code line-by-line which means if there is a function that causes the
routine to hang (like http.Get()) the go routine will just pause and do nothing.

One CPU core runs a Go Scheduler which schedules GRs. Scheduler will run one
routine until it finishes or makes a blocking call (like an HTTP request). Go tries to use
one core by default. If we have multiple CPU cores, the scheduler can run a GR on each core.
Scheduler runs one thread on each "logical" core.

Channels are typed (if a channel is type string, it cannot share non-string data)
*/
func main() {
	//channelExample()
	//waitGroupExample()

	fmt.Println("Multiple write to channel")
	c := make(chan string)

	var wg sync.WaitGroup

	wg.Add(1)
	go channelExperiment(c, &wg)
	wg.Wait()
}

func channelExperiment(c chan string, wg *sync.WaitGroup) {
	c <- sendDataA()
	fmt.Println(<-c)

	c <- sendDataB()
	fmt.Println(<-c)

	wg.Done()
}

func sendDataA() string {
	return "A"
}

func sendDataB() string {
	return "B"
}

func waitGroupExample() {
	fmt.Println("Go WaitGroup tutorial")
	var wg sync.WaitGroup

	wg.Add(1)
	go myFunc(&wg)
	wg.Wait()

	fmt.Println("Finished executing my go program")
}

func myFunc(wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second)
	fmt.Println("Finished Executing Goroutine")
	wg.Done()
}

func channelExample() {
	links := []string{
		"http://google.com",
		"http://reddit.com",
		"http://facebook.com",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		// checkLink(link) without go routine, will cause main go routine to hang

		// with go routine. Whenever a go routine hangs, the main GR will kick off a new child GR
		go checkLink(link, c)
	}

	/*
		for i := 0; i < len(links); i++ {
			go checkLink(<-c, c)
			//fmt.Println(<- c) // Blocking channel. Main GR will wait until channel gets data
		}
	*/

	// Check status until eternity
	for {
		go checkLink(<-c, c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
