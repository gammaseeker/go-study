package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	for i := range d {
		newPosition := rng.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}

}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// Option 1 - log the error and return a call to newDeck()
		// Option 2 - log the error and entirely quit the program
		// We arbitrarily go with option 2
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	deckString := string(bs) // Convert []byte to a string
	s := strings.Split(deckString, ",")
	return deck(s)
}
