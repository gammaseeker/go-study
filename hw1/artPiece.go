package main

import "fmt"

// for linkedlist you need structs of nodes chained together and another
// struct to contain the entire collection. The wrapper struct should
// have head of the chain of nodes
type artPiece struct {
	id         int
	artType    string
	artName    string
	artistName string
	price      int
	next       *artPiece
}

type artPieces struct {
	head *artPiece
	len  int
}

func (a *artPieces) PrintList() {
	head := a.head
	for head != nil {
		fmt.Printf("%+v\n", head)
		head = head.next
	}
}

func (a *artPieces) Insert(id int, artType string,
	artName string, artistName string, price int) {
	newArt := artPiece{
		id:         id,
		artType:    artType,
		artName:    artName,
		artistName: artistName,
		price:      price,
		next:       nil,
	}

	if a.len == 0 {
		a.head = &newArt
		a.len++
		return
	}

	curr := a.head
	for i := 0; i < a.len; i++ {
		if curr.next == nil {
			curr.next = &newArt
			a.len++
			return
		}
		curr = curr.next
	}
}
