package main

import "fmt"

func main() {
	donuts := map[string]int{
		"frosted":   10,
		"chocolate": 15,
		"jelly":     8,
	}
	fmt.Println(donuts)

	delete(donuts, "frosted")
	delete(donuts, "chocolate")
	delete(donuts, "jelly")

	fmt.Println(donuts)
}
