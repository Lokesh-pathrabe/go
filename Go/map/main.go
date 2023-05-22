package main

import "fmt"

func main() {

	// var colors map[string]string

	// colors := make(map[string]string)
	colors1 := make(map[int]string)

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}
	colors["white"] = "#dfdfdf"
	colors1[10] = "#dfdfdf"

	// fmt.Println(colors)
	// fmt.Println(colors1)

	delete(colors1, 10)
	fmt.Println(colors1)
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code of", color, ":", hex)
	}
}
