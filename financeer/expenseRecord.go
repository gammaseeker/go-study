package main

import "fmt"

type expenseRecord struct {
	expenseName string
	category    string
	price       float32
}

func newExpenseRecord(expenseName string, category string, price float32) expenseRecord {
	expense := expenseRecord{
		expenseName: expenseName,
		category:    category,
		price:       price,
	}
	return expense
}

func (e expenseRecord) printExpense() {
	fmt.Printf("%+v", e)
}
