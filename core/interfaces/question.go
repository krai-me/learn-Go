// solve the errs in it
package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	peri() int
}
type triangle struct {
	base, height float32
}

func (t triangle) area() float64 {
	return 0.5 * t.base * t.height
}
func (t triangle) perim() float64 {
	return t.base + t.height + math.Sqrt((math.Pow(t.base, 2) + math.Pow(t.height, 2)))
}

func method(g geometry) {
	fmt.Println("the area is:", g.area())
	fmt.Println("the peri is:", g.perim())

}

func main() {
	tridata := triangle{base: 10, height: 10}
	method(tridata)
}
