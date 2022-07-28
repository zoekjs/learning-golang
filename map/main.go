package main

import "fmt"

func main() {
	// var colors map[string]string -> declare a map without values

	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#00ff00",
	// }

	colors := make(map[string]string)

	//add value to map
	colors["white"] = "#ffffff"

	fmt.Println(colors)
}
