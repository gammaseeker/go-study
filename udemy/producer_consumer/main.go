package main

import (
	"math/rand"
	"time"

	"github.com/fatih/color"
)

const NumberOfPizzas = 10

var pizzasMade, pizzasFailed, total int

type Producer struct {
	data chan PizzaOrder
	quit chan chan error // A channel of channels
}

type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}

func pizzeria(pizzaMaker *Producer) {
	// keep track of which pizza we are making

	// run forever or until we receive a quit notif
	// try to make pizzas
	for {
		// Try to make a pizza
		// decision
	}
}

func main() {
	// seed the random num generator
	rand.Seed(time.Now().UnixNano())

	// print out msg
	color.Cyan("The Pizzeria is open for business!")
	color.Cyan("----------------------------------")

	// create a producer
	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}
	// run the producer in the background
	go pizzeria(pizzaJob)

	// create and run consumer
	// print out the ending message
}
