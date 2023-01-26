package main

import "fmt"

// Cannot directly create a value of an interface. It is not a concrete type. It is not a generic type
// Interfaces are 'implicit' don't need to explicitly say which type belongs to an interface
type bot interface {
	getGreeting() string // function signature should match where it is declared
	// Can also nest interfaces within each other. Look at the ReadCloser interface on go docs
}

type englishBot struct{}
type koreanBot struct{}

func main() {
	eb := englishBot{}
	kb := koreanBot{}

	printGreeting(eb)
	printGreeting(kb)
}

func printGreeting(b bot) {
	// If you are a type with a getGreeting(), you are an honorary member of type 'bot'
	fmt.Println(b.getGreeting())
}

/* Golang does not support overloading, these two functions are similar in nature
func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

func printGreeting(kb koreanBot) {
	fmt.Println(kb.getGreeting())
}
*/

// Since we are not using receiver value, we omit the receiver variable
func (englishBot) getGreeting() string {
	// VERY cutsom logic for generating an english greeting
	return "Hi there!"
}

func (koreanBot) getGreeting() string {
	// VERY cutsom logic for generating a korean greeting
	return "안녕하세요!"
}
