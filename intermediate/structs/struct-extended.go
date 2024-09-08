package main

import "fmt"

// Define the Rectangle struct
type Rectangle struct {
	length float32
	width  float32
}

// Method to calculate the area of the rectangle
// (r Rectangle) is the receiver for the method area.
// Inside the method, r refers to the Rectangle instance on which the method is called.
func (r Rectangle) area() float32 {
	return r.length * r.width
}

func main() {
	// Create an instance of Rectangle
	rect := Rectangle{length: 5.0, width: 3.0}

	// Call the area() method on the Rectangle instance
	fmt.Println("Area of the rectangle:", rect.area()) // Output: Area of the rectangle: 15
}
