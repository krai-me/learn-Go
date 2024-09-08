package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius int
}

func (c Circle) area() float32 {
	return float32(c.radius) * math.Pi * float32(c.radius)
}
func (c Circle) circumference() float32 {
	return float32(c.radius) * 2 * math.Pi
}

func main() {
	var input int
	fmt.Println("Enter the radius")
	fmt.Scan(&input)
	rad := Circle{radius: input}
	fmt.Println("the circumference is:", rad.area())
	fmt.Println("the circumference is:", rad.circumference())
}
