package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

type Income struct {
	Source string
	Amount int
}

func updateMessage(s string, m *sync.Mutex) {
	defer wg.Done()

	m.Lock()
	msg = s
	m.Unlock()
}

// go run -race <go files> is a diagnostic tool to detect race conditions
func main() {
	/*
		msg = "Hello, world!"

		var mutex sync.Mutex // Never copy mutexes once it is created

		// Race condition. Need mutex to fix
		wg.Add(2)
		go updateMessage("Hello, universe!", &mutex)
		go updateMessage("Hello, cosmos!", &mutex)
		wg.Wait()

		fmt.Println(msg)
	*/

	// variable for bank balance
	var bankBalance int
	var balance sync.Mutex

	// print out starting values
	fmt.Printf("Initial account balance: $%d.00\n", bankBalance)

	// define weekly revenue
	incomes := []Income{
		{Source: "Main job", Amount: 500},
		{Source: "Gifts", Amount: 10},
		{Source: "Part time job", Amount: 50},
		{Source: "Investments", Amount: 100},
	}

	wg.Add((len(incomes)))

	// loop through 52 weeks and print out how much is made; keep a running total
	for i, income := range incomes {
		go func(i int, income Income) {
			defer wg.Done()

			for week := 1; week <= 52; week++ {
				balance.Lock()
				temp := bankBalance
				temp += income.Amount
				bankBalance = temp
				balance.Unlock()

				fmt.Printf("On week %d, you earned $%d.00 from %s\n", week, income.Amount, income.Source)
			}
		}(i, income)
	}

	wg.Wait()

	// print out final balance
	fmt.Printf("FInal bank balance: $%d.00\n", bankBalance)
}
