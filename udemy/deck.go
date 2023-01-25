package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func getSuits() [4]string {
	return [4]string{"Spades", "Hearts", "Diamonds", "Clubs"}
}

func getValues() [4]string {
	return [4]string{"Ace", "Two", "Three", "Four"}
}

func newDeck() deck {
	cards := deck{}

	cardSuits := getSuits()
	cardValues := getValues()

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			card := value + " of " + suit
			cards = append(cards, card)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	// []string(d) converts from type deck to string slice
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func (d deck) shuffle() {

}

func newDeckFromFile() {

}
