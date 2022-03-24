package main

import "fmt"

func main() {

	// var colors map[string]string

	//colors := make(map[string]string)

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}

	colors["white"] = "#ffffff"

	printMap(colors)

}

func printMap(c map[string]string) {
	for k, v := range c {
		fmt.Printf("%+v is %+v\n", k, v)
	}
}
