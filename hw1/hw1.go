package main

import (
	"fmt"
	"os"
)

// use case: ./hw1 filename -b N [-v] [-i id] [-t type] [-n artist_name] [-o filename]
func main() {
	node3 := artPiece{
		id:         3,
		artType:    "a",
		artName:    "b",
		artistName: "c",
		price:      1,
		next:       nil,
	}
	node2 := artPiece{
		id:         2,
		artType:    "a",
		artName:    "b",
		artistName: "c",
		price:      1,
		next:       &node3,
	}
	node1 := artPiece{
		id:         1,
		artType:    "a",
		artName:    "b",
		artistName: "c",
		price:      1,
		next:       &node2,
	}

	list := artPieces{
		head: &node1,
	}
	PrintList(&list)
	return
	argsList := os.Args
	numArgs := len(argsList)

	// Min number of args will be 5
	if numArgs < 5 {
		fmt.Println("Error: must have minimum of 5 arguments\n NO QUERY PROVIDED")
		return
	}

	filename := argsList[1]
	budget := argsList[3]

	fmt.Println(filename)
	fmt.Println(budget)

	var vFlag bool = false
	if numArgs > 4 {
		for i := 4; i < numArgs; i++ {
			switch argsList[i] {
			case "-v":
				fmt.Println(argsList[i])
				vFlag = true
			case "-i":
				if !vFlag {
					id := argsList[i+1]
					fmt.Println(id)
				}
			case "-t":
				if !vFlag {
					artType := argsList[i+1]
					fmt.Println(artType)
				}
			case "-n":
				if !vFlag {
					artistName := argsList[i+1]
					fmt.Println(artistName)
				}
			case "-o":
				outFilename := argsList[i+1]
				fmt.Println(outFilename)
			}
		}
	}

}
