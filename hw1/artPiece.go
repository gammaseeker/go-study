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
}

func PrintList(list *artPieces) {
	head := list.head
	for head != nil {
		fmt.Printf("%+v\n", head)
		head = head.next
	}
}
