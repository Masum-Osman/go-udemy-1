package main2

import "fmt"

func main() {

	/**
	Diff ways of creating a map:

	1. var colors map[string]string
	2. colors := make(map[string]string)
	3. 	colors := map[string]string{
		"red":   "ff0000",
		"green": "4bf745",
	}

	*/
	colors := map[string]string{
		"red":   "ff0000",
		"green": "4bf745",
		"white": "ffffff",
	}

	// delete(colors, "red")

	// fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, value := range c {
		fmt.Println(color, value)
	}
}
