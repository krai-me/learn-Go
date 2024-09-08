package main

import "fmt"

// Define the Point struct
// Each field within a struct has a name and a type, allowing you to encapsulate and manage related data easily.
type Point struct {
	x int // The x coordinate of the point in a 2D space
	y int // The y coordinate of the point in a 2D space
}

func main() {
	// Create an instance of the Point struct
	// Using shorthand declaration and initialization
	p := Point{x: 10, y: 20}
	// or alternatively
	// var p Point
	// p = Point{x: 10, y: 20}

	fmt.Println("Point instance:", p) // Output: Point instance: {10 20}

	// Also valid: Initializes only the x coordinate; y will be set to its zero value (0)
	p = Point{x: 12}
	fmt.Println("Point instance with only x set:", p) // Output: Point instance with only x set: {12 0}

	// Resets the Point instance to its zero value
	p = Point{}                              // x and y will be set to 0 by default
	fmt.Println("Point after resetting:", p) // Output: Point after resetting: {0 0}

	// Accessing struct fields
	fmt.Println("X coordinate of Point instance:", p.x) // Output: X coordinate of Point instance: 0

	// Updating struct field
	p.x = 10
	p.y = 15
	fmt.Println("Point after updating fields:", p) // Output: Point after updating fields: {10 15}
}
