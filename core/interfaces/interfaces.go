// An interface in Go is a type that specifies a contract:
// a collection of method signatures that a type must implement.
// It doesn’t define how these methods are implemented, only that they exist.
// In simple words interfaces in Go help manage multiple methods and types in a very flexible and organized way.
package main

import (
	"fmt"
	"math"
)

// Define an interface named geometry
// An interface specifies a contract with method signatures that types must implement.
type geometry interface {
	area() float64  // Method to calculate the area
	perim() float64 // Method to calculate the perimeter
}

// Define a rect type
type rect struct {
	width, height float64
}

// Implement the geometry interface for rect
func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// Define a circle type
type circle struct {
	radius float64
}

// Implement the geometry interface for circle
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// Function to measure the geometry of any type that implements the geometry interface
func measure(g geometry) {
	fmt.Println("Shape details:")
	fmt.Println("Area:", g.area())
	fmt.Println("Perimeter:", g.perim())
}

func main() {
	// Create instances of rect and circle
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	// Use the measure function to print details of the shapes
	measure(r)
	measure(c)
}
