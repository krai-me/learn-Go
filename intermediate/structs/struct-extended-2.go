package main

import "fmt"

// Define a Point struct to represent a point in 2D space
type Point struct {
	x int
	y int
}

func main() {
	// Create a slice of Points with specific coordinates
	points := []Point{{x: 1, y: 1}, {x: 7, y: 27}, {x: 12, y: 7}, {x: 9, y: 25}}

	// Print the first point in the slice
	fmt.Println("First point:", points[0]) // Output: First point: {1 1}

	// Modify the second point in the slice
	points[1].x = 8
	points[1].y = 16

	// Print the updated second point
	fmt.Println("Updated second point:", points[1]) // Output: Updated second point: {8 16}

	// You can also create named points and add them to a slice
	a := Point{1, 1}
	b := Point{7, 27}
	c := Point{12, 7}
	d := Point{9, 25}

	namedPoints := []Point{a, b, c, d}

	for i, point := range namedPoints {
		fmt.Println("Point", i+1, ":", point)
	}
}
