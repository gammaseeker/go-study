package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

var msg string

func updateMessage(s string) {
	msg = s
}

func printMessage() {
	fmt.Println(msg)
}

func updateAndPrintMessage(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	updateMessage(s)
	printMessage()
}

func main() {

	var wg sync.WaitGroup

	// challenge: modify this code so that the calls to updateMessage() on lines
	// 28, 30, and 33 run as goroutines, and implement wait groups so that
	// the program runs properly, and prints out three different messages.
	// Then, write a test for all three functions in this program: updateMessage(),
	// printMessage(), and main().

	msg = "Hello, world!"

	wg.Add(2)
	go updateAndPrintMessage("Hello, universe!", &wg)
	go updateAndPrintMessage("Hello, cosmos!", &wg)
	wg.Wait()

	updateMessage("Hello, world!")

	printMessage()
}

func Test_updateMessage(t *testing.T) {
	testString := "test"

	updateMessage(testString)
	if msg != testString {
		t.Errorf("msg does not match the testString")
	}
}

func Test_printMessage(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	msg = "test"
	printMessage()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "test") {
		t.Errorf("Exepected, to find test, but it is not there")
	}
}
