package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

type Shape interface {
	area() float64
	perimeter() float64
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x2, r.y1)
	h := distance(r.x1, r.y1, r.x1, r.y2)
	return l * h
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (r *Rectangle) perimeter() float64 {
	l := math.Abs(r.x2 - r.x1)
	h := math.Abs(r.x2 - r.x1)
	return 2 * (l + h)
}

func (c *Circle) perimeter() float64 {
	return math.Pi * 2 * c.r
}

func main() {
	c := Circle{1, 2, 1}
	fmt.Printf("circle area: %v\n", c.area())
	fmt.Printf("circle perimeter: %v\n", c.perimeter())
	r := Rectangle{1, 1, 3, 4}
	fmt.Printf("rectangle area: %v\n", r.area())
	fmt.Printf("rectangle perimeter: %v\n", r.perimeter())
}
