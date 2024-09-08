// A method in Go is a function that is associated with a specific type (usually a struct).
// It allows you to operate on the data within that type.
// In simple words Methods allow you to attach functions to types.
package main

import "fmt"

type rect struct {
	width, height int
}

// Method with Pointer Receiver:
// The receiver r *rect is a pointer to a rect struct,
// allowing the method to operate on the original struct and avoid copying.
func (r *rect) area() int {
	return r.width * r.height
}

// Method with Value Receiver:
// The receiver r rect is a value type, meaning a copy of the struct is used.
// Modifications within this method do not affect the original struct.

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	// Call methods on the value receiver
	fmt.Println("area: ", r.area())  // Output: area: 50
	fmt.Println("perim:", r.perim()) // Output: perim: 30

	// Call methods on the pointer receiver
	rp := &r
	fmt.Println("area: ", rp.area())  // Output: area: 50
	fmt.Println("perim:", rp.perim()) // Output: perim: 30
}
