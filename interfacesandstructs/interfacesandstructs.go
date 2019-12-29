package interfacesandstructs

import "math"

// Shape interface
type Shape interface {
	Area() float64
}

// Rectangle struct
type Rectangle struct {
	Height float64
	Width  float64
}

// Circle struct
type Circle struct {
	Radius float64
}

// Triangle struct
type Triangle struct {
	Base   float64
	Height float64
}

// Area method for Circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perimeter method for Rectangle
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Height + r.Width)
}

// Area method for Rectangle
func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

// Area method for Triangle
func (t Triangle) Area() float64 {
	return t.Height * t.Base / 2
}
