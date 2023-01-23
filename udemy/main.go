package main

import "fmt"

func main() {
	//var card string = "Ace of Spades"
	card := newCard()
	//card = "Five of Diamonds"
	//cards := []string{"Ace of Diamonds", newCard()} // Slices are static
	cards := deck{"Ace of Diamonds", newCard()} // Slices are static
	cards = append(cards, "Six of Spades")      // append() returns a new slice

	/*
		// i is index of this element in array
		// c is current element we're iterating over
		// range cards is collection we're looping over
		for i, c := range cards {
			fmt.Println(i, c)
		}
	*/

	cards.print()
	fmt.Println(card)
	fmt.Println(cards)

	myDeck := newDeck()
	hand, remainingDeck := deal(myDeck, 2)
	hand.print()
	remainingDeck.print()
}

func newCard() string {
	return "Five of Diamonds"
}
