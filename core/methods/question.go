//fix err in code even in comments
package main

import "fmt"

// Define the rect struct
type rect struct {
	width  float64
	height int
}

// Method to double the size of the rectangle dimensions (pointer receiver)
func (r rect) doubleSize() (float64, float64) {
	r.width *= 2
	r.height *= 2
	return r.width, r.height
}

// Method to display the current dimensions of the rectangle (value receiver)
func (r *rect) displayDimensions() (float64, int) {
	return r.width, r.height
}

func main() {
	// Incorrect initialization of the rect struct
	r := rect{width: "10", height: 10}

	// Display original dimensions
	fmt.Println("Original dimensions:", r.displayDimensions())

	// Double the size of the rectangle
	fmt.Println(r.doubleSize())

	// Display new dimensions after doubling
	fmt.Println("New dimensions after doubling:", r.displayDimensions())
}
