package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Name() string
}

type triangle struct {
	base   float64
	height float64
}
type rectangle struct {
	length float64
	width  float64
}
type circle struct {
	radius float64
}

func (t triangle) Area() float64 {
	return (t.base * t.height) / 2
}

func (t triangle) Name() string {
	return "triangle"
}

func (r rectangle) Area() float64 {
	return r.length * r.width
}

func (r rectangle) Name() string {
	return "rectangle"
}

func (c circle) Area() float64 {
	return 2.0 * math.Pi * c.radius
}

func (c circle) Name() string {
	return "circle"
}

func printShapeDetails(shapes ...Shape) {
	for _, item := range shapes {
		fmt.Printf("The area of %s is: %.2f\n", item.Name(), item.Area())
	}
}

func main() {
	t := triangle{base: 15.5, height: 20.1}
	r := rectangle{length: 20, width: 10}
	c := circle{radius: 10}
	printShapeDetails(t, r, c)
}
