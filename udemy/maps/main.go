package main

import "fmt"

func main() {
	// Ways of declaring a map

	//var colors map[string]string produces an empty map
	// colors := make(map[string]string)

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}

	// To mutate map
	colors["white"] = "#ffffff"
	//delete(colors, "white")

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
