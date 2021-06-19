package main

import "fmt"

func main() {
	colors := mapOfColors()

	printMap(colors)
}

func mapOfColors() map[string]string {
	colors := map[string]string{
		"red":    "#fd977",
		"yellow": "#87fd98",
		"white":  "#ffffff",
	}
	return colors
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for color", color, "is", hex)
	}
}
